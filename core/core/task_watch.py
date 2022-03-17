#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-11 16:45:45
LastEditors: zmf96
LastEditTime: 2022-03-07 08:28:05
FilePath: /task_watch.py
Description: 
'''
from concurrent.futures import ThreadPoolExecutor
import json
import time

import pika

import tasks
from common.config import broker_url
from common.log import logger
from consum_data_done import consum_data_done
from model import Task

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:default', durable=False)

threadPoolTaskWatch = ThreadPoolExecutor(
    max_workers=16, thread_name_prefix="task_watch_")

# 6. 消费任务结果


def consum_task_date(celery_task_ids, task):
    for task_id in celery_task_ids:
        try:
            res = tasks.app.AsyncResult(task_id).get()
            logger.info(res)
            consum_data_done(res.get("tool_type"),
                             res.get("data"), task)
        except Exception as e:
            logger.warning(e)

# 5. 向celery worker节点，分发任务


def task_worker_run(tools, task):
    celery_task_ids = []
    # TODO: 优化代码！
    for tool in tools:
        if tool == "gettitle":
            for host in task.get("hosts").split(","):
                for port in task.get("ports").split(","):
                    if "443" in port:
                        res = tasks.gettitle.delay("https://"+host+":"+port)
                    else:
                        res = tasks.gettitle.delay("http://"+host+":"+port)
                    logger.info(res)
                    celery_task_ids.append(res.id)
        elif tool == "beian2domain":
            for keyword in task.get("keyword").split(","):
                res = tasks.beian2domain.delay(keyword)
                logger.info(res)
                celery_task_ids.append(res.id)
        elif tool == "cdncheck":
            for host in task.get("hosts").split(","):
                res = tasks.cdncheck.delay(host)
                logger.info(res)
                celery_task_ids.append(res.id)
        elif tool == "pysubdomain":
            for host in task.get("hosts").split(","):
                res = tasks.pysubdomain.delay(host)
                logger.info(res)
                celery_task_ids.append(res.id)
        elif tool == "hotfinger":
            port_str = task.get("ports").strip()
            if len(port_str) > 0:
                port_list = port_str.split(",")
            else:
                port_list = ["80", "443"]
            print(port_list)
            for host in task.get("hosts").split(","):
                for port in port_list:
                    if "443" in port:
                        res = tasks.hotfinger.delay("https://"+host+":"+port)
                    else:
                        res = tasks.hotfinger.delay("http://"+host+":"+port)
                    logger.info(res)
                    celery_task_ids.append(res.id)
        elif tool == "fofainfo":
            res = tasks.fofainfo.delay(task.get("keyword"))
            logger.info(res)
            celery_task_ids.append(res.id)
        elif tool == "emailall":
            for host in task.get("hosts").split(","):
                logger.info(host)
                res = tasks.emailall.delay(host)
                logger.info(res)
                celery_task_ids.append(res.id)
        else:
            logger.warning("Not support tool: %s" % tool)

    consum_task_date(celery_task_ids, task)
    return celery_task_ids

# 4. tool分组


def task_dely(tools, task):
    tools = set(tools)
    celery_task_ids = []
    tools_group_sort = [
        {"pysubdomain", "beian2domain"},
        {"fofainfo"},
        {"gettitle", "cdncheck", "hotfinger"},
        {"emailall"}
    ]
    current_tools = []
    for tool_group in tools_group_sort:
        tmp = tools & tool_group
        if len(tmp) > 0:
            current_tools.append(tmp)
    celery_task_ids = []
    for current_tool in current_tools:
        current_ids = task_worker_run(current_tool, task)
        celery_task_ids.extend(current_ids)

    return celery_task_ids

# 3. 添加工作worker


def worker(task):
    celery_task_ids = []
    tools = task.get("tools", [])
    if isinstance(tools, list):
        task_obj = Task.objects(id=task.get("_id")).first()
        task_obj.Update(status="running")
        celery_task_ids = task_dely(tools, task)
        logger.info("celery_task_ids: %s" % celery_task_ids)
        task_obj.Update(status="complete",
                        celery_task_ids=task.get("celery_task_ids"))

# 2. 消费task消息


def callback(ch, method, properties, body):
    logger.info(body)
    if properties.type == "task":
        task = json.loads(body.decode("utf-8"))
        threadPoolTaskWatch.submit(worker, task)
    else:
        print(" [x] Received %r" % body)
    ch.basic_ack(delivery_tag=method.delivery_tag)  # 告诉生产者，消息处理完成

# 1. 启动


def start_done():
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume("server:default", callback, False)
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


if __name__ == "__main__":
    start_done()
