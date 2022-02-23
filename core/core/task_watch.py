#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-11 16:45:45
LastEditors: zmf96
LastEditTime: 2022-02-23 08:14:22
FilePath: /core/core/task_watch.py
Description: 
'''
from concurrent.futures import ThreadPoolExecutor
import json
import queue

import pika

import tasks
from common.config import broker_url
from common.log import logger
from consum_data_done import consum_data_done
from model import Task

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:default', durable=False)

push_channel = connection.channel()
push_channel.queue_declare(queue='server:task_done', durable=False)

threadPool = ThreadPoolExecutor(
    max_workers=10, thread_name_prefix="task_watch_")

push_queue = queue.Queue()

def push_message():
    global push_channel
    while True:
        task = push_queue.get()
        push_channel.basic_publish(exchange='', routing_key="server:task_done", body=json.dumps(
            task), properties=pika.BasicProperties(type="task_done"))


def task_worker_run(tools, task, wait_done=False):
    celery_task_ids = []
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
        else:
            logger.warning("Not support tool: %s" % tool)

    if wait_done:
        for task_id in celery_task_ids:
            try:
                res = tasks.app.AsyncResult(task_id).get()
                logger.info(res)
                task_obj = Task.objects(id=task.get("_id")).first()
                consum_data_done(res.get("tool_type"),
                                 res.get("data"), task_obj)
            except Exception as e:
                logger.warning(e)
    return celery_task_ids


def task_dely(tools, task):
    tools = set(tools)
    celery_task_ids = []
    tools_group_sort = [
        {"pysubdomain", "beian2domain"},
        {"gettitle", "cdncheck"},
    ]
    current_tools = []
    for tool_group in tools_group_sort:
        tmp = tools & tool_group
        if len(tmp) > 0:
            current_tools.append(tmp)
    if len(current_tools) > 1:
        for current_tool in current_tools[:-1]:
            task_worker_run(current_tool, task, wait_done=True)
        celery_task_ids = task_worker_run(current_tools[-1], task)
    elif len(current_tools) == 1:
        task_worker_run(current_tools[0], task)
    return celery_task_ids


def worker(task):
    celery_task_ids = []
    tools = task.get("tools", [])
    if isinstance(tools, list):
        celery_task_ids = task_dely(tools, task)
        logger.info("celery_task_ids: %s" % celery_task_ids)
        task["status"] = "running"
        task["celery_task_ids"] = celery_task_ids
    push_queue.put(task)



def callback(ch, method, properties, body):
    logger.info(body)
    if properties.type == "task":
        task = json.loads(body.decode("utf-8"))
        threadPool.submit(worker, task)
    else:
        print(" [x] Received %r" % body)
    ch.basic_ack(delivery_tag=method.delivery_tag)  # 告诉生产者，消息处理完成


def start_done():
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(on_message_callback=callback,
                          queue="server:default")
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


if __name__ == "__main__":
    try:
        start_done()
    except Exception as e:
        logger.error(e)
