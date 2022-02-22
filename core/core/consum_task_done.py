#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-15 14:36:13
LastEditors: zmf96
LastEditTime: 2022-02-22 07:20:33
FilePath: /core/core/consum_task_done.py
Description: 
'''

import json
from concurrent.futures import ThreadPoolExecutor

import pika
from tasks import app
from common.config import broker_url
from common.log import logger
from model import Task,PathInfo,Domain,PortInfo

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:task_done', durable=False)

threadPool = ThreadPoolExecutor(
    max_workers=10, thread_name_prefix="task_done_")

def consum_gettile(data,task_obj):
    try:
        tmp = data.get("url").split("://")[1].split("/")
        data["hostinfo"] = tmp[0]
        data["path"] = "/"+"/".join(tmp[1:])
        data["target_id"] = task_obj.target_id
        data["project_id"] = task_obj.project_id
        data["response_size"] = len(data.get("body"))
        if PathInfo.objects(url=data.get("url")).count() > 0:
            PathInfo.objects(url=data.get("url")).first().Update(**data)
        else:
            PathInfo(**data).Save()
        if data["path"] == "/":
            data.pop("path")
            data.pop("response_size")
            url = data.pop("url")
            tmp = data.get("hostinfo").split(":")
            data["host"] = tmp[0]
            if len(tmp) > 1:
                port = tmp[1]
            elif url.startswith("https://"):
                port = "443"
            else:
                port = "80"
            data["port"] = port
            if PortInfo.objects(hostinfo=data.get("hostinfo")).count() > 0:
                PortInfo.objects(hostinfo=data.get("hostinfo")).first().Update(**data)
            else:
                PortInfo(**data).Save()
    except Exception as e:
        logger.warning(e)

def consum_cdncheck(data,task_obj):
    for item in data:
        logger.info(item)
        item["target_id"] = task_obj.target_id
        item["project_id"] = task_obj.project_id
        cdn = 0
        if item.get("cdn") != '':
            cname = item.get("cdn")
            cdn = 1
        if Domain.objects(domain=item.get("target")).count() > 0:
            do = Domain.objects(domain=item.get("target")).first()
            if item.get("ip") not in do.ips:
                do.ips.append(item.get("ip"))
            do.cdn = cdn
            do.cname = cname
            do.Save()
        else:
            Domain(domain=item.get("target"),cdn=cdn,cname=cname,ips=[item.get("ip")]).Save()

def consum_beian2domain(data,task_obj):
    for item in data:
        logger.info(item)
        tmp = {"domain":item[1],"target_id":task_obj.target_id,"project_id":task_obj.project_id,"tags":[item[0]]}
        if Domain.objects(domain=item[1]).count() > 0:
            do = Domain.objects(domain=item[1]).first()
            if item[0] not in do.tags:
                do.tags.append(item[0])
                do.Save()
        else:
            Domain(**tmp).Save()

def consum_pysubdomain(data,task_obj):
    for k,v in data.items():
        logger.info(k,v)
        tmp = {"domain":k,"target_id":task_obj.target_id,"project_id":task_obj.project_id,"ips":v}
        if Domain.objects(domain=k).count() > 0:
            # do = Domain.objects(domain=k).first()
            pass
        else:
            Domain(**tmp).Save()

def consum_done(task):
    try:
        logger.info(task)
        task_obj = Task.objects(id=task.get("_id")).first()
        task_obj.Update(status="running",celery_task_ids=task.get("celery_task_ids"))
        for task_id in task.get("celery_task_ids"):
            logger.info(task_id)
            res = app.AsyncResult(task_id).get()
            logger.info(res)
            if res.get("tool_type") == "gettitle":
                tmp_data = res["data"][0]
                consum_gettile(tmp_data,task_obj)
            elif res.get("tool_type") == "cdncheck":
                consum_cdncheck(res["data"],task_obj)
            elif res.get("tool_type") == "beian2domain":
                consum_beian2domain(res.get("data"),task_obj)
            elif res.get("tool_type") == "pysubdomain":
                consum_pysubdomain(res.get("data"),task_obj)
            else:
                pass
        task_obj.Update(status="complete")

    except Exception as e:
        logger.warning(e)
    logger.info("done")



def callback(ch, method, properties, body):
    logger.info(body)
    if properties.type == "task_done":
        # todo: running
        task = json.loads(body.decode("utf-8"))
        threadPool.submit(consum_done, task)
    else:
        print(" [x] Received %r" % body)
    ch.basic_ack(delivery_tag=method.delivery_tag)


def start_done():
    global channel
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(on_message_callback=callback,
                          queue="server:task_done")
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


if __name__ == "__main__":
    try:
        start_done()
    except Exception as e:
        logger.error(e)
