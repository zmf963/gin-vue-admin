#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-11 16:45:45
LastEditors: zmf96
LastEditTime: 2022-03-29 19:20:44
FilePath: /core/core/task_monitoring.py
Description: 
"""
from concurrent.futures import ThreadPoolExecutor
from typing import Set
import json
import time
from globals import MODEL_OBJECT

import pika

import tasks
from config import broker_url
from cp_common.log import logger
from cp_common.utils import get_url_list_from_host_port_list, parse_host

# from consum_data_done import consum_data_done
from model.model import Task, Domain

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue="server:default", durable=False)

threadPoolTaskWatch = ThreadPoolExecutor(
    max_workers=16, thread_name_prefix="task_watch_"
)

# 6. 消费任务结果


def consum_task_data(celery_task_ids, task):
    for task_id in celery_task_ids:
        try:
            logger.warning(task_id)
            result_list = tasks.app.AsyncResult(task_id).get()
            logger.info(task)
            for result in result_list:
                logger.info(result)
                for k, v in result.items():
                    if k in MODEL_OBJECT:
                        v["target_id"] = task.get("target_id")
                        MODEL_OBJECT[k].Add(v)
        except Exception as e:
            logger.warning(e)


def get_tool_url_set(task: object, tool_name: str) -> Set[str]:
    host_list = parse_host(task.get("hosts"))
    port_list = parse_host(task.get("ports"))
    _tool_ext = task.get("tool_ext")
    try:
        if _tool_ext[tool_name]["host"] == "query":
            url_list = set()
            for host in host_list:
                for dom in Domain.objects(domain__endswith=host):
                    url_list = url_list | get_url_list_from_host_port_list(
                        [dom.domain], port_list
                    )
            return url_list
    except Exception as e:
        logger.error(e)
    return get_url_list_from_host_port_list(host_list, port_list)


# 5. 向celery worker节点，分发任务


def task_worker_run(tools, task):
    celery_task_ids = []
    for tool in tools:
        try:
            url_list = task.get("tool_ext").get("url_list", [])
            if len(url_list) < 1:
                url_list = list(get_tool_url_set(task, tool))
            res = tasks.worker_entry.delay(
                tool,
                host_list=task.get("hosts"),
                port_list=task.get("ports"),
                keyword=task.get("keyword"),
                url_list=url_list,
                **task.get("tool_ext")
            )
            logger.info(res)
            celery_task_ids.append(res.id)
        except Exception as e:
            logger.warning(e)
            logger.warning("Not support tool: %s" % tool)

    consum_task_data(celery_task_ids, task)
    return celery_task_ids


# 4. tool分组


def task_dely(tools, task):
    tools = set(tools)
    celery_task_ids = []
    tools_group_sort = [
        {"pysubdomain", "beian2domain"},
        {"fofainfo"},
        {"gettitle", "cdncheck", "hotfinger"},
        {"emailall"},
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
        task_obj.Update(status="complete", celery_task_ids=task.get("celery_task_ids"))


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
    print(" [*] Waiting for messages. To exit press CTRL+C")
    channel.start_consuming()


if __name__ == "__main__":
    start_done()
