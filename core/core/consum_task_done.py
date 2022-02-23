#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-15 14:36:13
LastEditors: zmf96
LastEditTime: 2022-02-23 03:41:24
FilePath: /core/core/consum_task_done.py
Description: 
'''

import json
from concurrent.futures import ThreadPoolExecutor

import pika

from common.config import broker_url
from common.log import logger
from consum_data_done import consum_data_done
from model import Task
from tasks import app

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:task_done', durable=False)

threadPool = ThreadPoolExecutor(
    max_workers=10, thread_name_prefix="task_done_")


def consum_done(task):
    try:
        logger.info(task)
        task_obj = Task.objects(id=task.get("_id")).first()
        task_obj.Update(status="running",
                        celery_task_ids=task.get("celery_task_ids"))
        for task_id in task.get("celery_task_ids"):
            logger.info(task_id)
            res = app.AsyncResult(task_id).get()
            logger.info(res)
            consum_data_done(res.get("tool_type"), res.get("data"), task_obj)
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
