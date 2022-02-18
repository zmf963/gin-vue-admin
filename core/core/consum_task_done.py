#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-15 14:36:13
LastEditors: zmf96
LastEditTime: 2022-02-18 03:29:23
FilePath: /core/consum_task_done.py
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
                try:
                    tmp = tmp_data.get("url").split("://")[1].split("/")
                    tmp_data["hostinfo"] = tmp[0]
                    tmp_data["path"] = "/"+"/".join(tmp[1:])
                    tmp_data["target_id"] = task_obj.target_id
                    tmp_data["project_id"] = task_obj.project_id
                    tmp_data["response_size"] = len(tmp_data.get("body"))
                    if PathInfo.objects(url=tmp_data.get("url")).count() > 0:
                        PathInfo.objects(url=tmp_data.get("url")).first().Update(**tmp_data)
                    else:
                        PathInfo(**tmp_data).Save()
                    if tmp_data["path"] == "/":
                        tmp_data.pop("path")
                        tmp_data.pop("response_size")
                        url = tmp_data.pop("url")
                        tmp = tmp_data.get("hostinfo").split(":")
                        tmp_data["host"] = tmp[0]
                        if len(tmp) > 1:
                            port = tmp[1]
                        elif url.startswith("https://"):
                            port = "443"
                        else:
                            port = "80"
                        tmp_data["port"] = port
                        if PortInfo.objects(hostinfo=tmp_data.get("hostinfo")).count() > 0:
                            PortInfo.objects(hostinfo=tmp_data.get("hostinfo")).first().Update(**tmp_data)
                        else:
                            PortInfo(**tmp_data).Save()
                except Exception as e:
                    logger.warning(e)

            elif res.get("tool_type") == "cdncheck":
                for item in res.get("data"):
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

            elif res.get("tool_type") == "beian2domain":
                for item in res.get("data"):
                    logger.info(item)
                    tmp = {"domain":item[1],"target_id":task_obj.target_id,"project_id":task_obj.project_id,"tags":[item[0]]}
                    if Domain.objects(domain=item[1]).count() > 0:
                        do = Domain.objects(domain=item[1]).first()
                        if item[0] not in do.tags:
                            do.tags.append(item[0])
                            do.Save()
                    else:
                        Domain(**tmp).Save()
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
