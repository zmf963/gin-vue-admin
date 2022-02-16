#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-11 16:45:45
LastEditors: zmf96
LastEditTime: 2022-02-15 15:14:10
FilePath: /core/task_watch.py
Description: 
'''
import tasks
import json
import pika
import time
from config import broker_url
from log import logger
connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:default', durable=False)

push_channel = connection.channel()
push_channel.queue_declare(queue='server:task_done', durable=False)


def push_message(task):
    logger.info(task)
    global push_channel
    push_channel.basic_publish(exchange='',routing_key="server:task_done",body=json.dumps(task),properties=pika.BasicProperties(type="task_done"))

def callback(ch, method, properties, body):
    logger.info(body)
    if properties.type == "task":
        task = json.loads(body.decode("utf-8"))
        celery_task_ids = []
        for tool in task.get("tools"):
            if tool == "gettitle":
                for host in task.get("hosts").split(","):
                    for port in task.get("ports").split(","):
                        res = tasks.gettitle.delay("http://"+host+":"+port)
                        logger.info(res)
                        celery_task_ids.append(res.id)
            else:
                logger.warning("Not support tool: %s" % tool)
        logger.info("celery_task_ids: %s" % celery_task_ids)
        task["status"] = "running"
        task["celery_task_ids"] = celery_task_ids
        push_message(task)
    else:
        print(" [x] Received %r" % body)
    ch.basic_ack(delivery_tag=method.delivery_tag)  # 告诉生产者，消息处理完成


channel.basic_qos(prefetch_count=1)  # 类似权重，按能力分发，如果有一个消息，就不在给你发
channel.basic_consume(on_message_callback=callback,
                      queue="server:default")

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
