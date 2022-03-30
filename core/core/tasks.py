#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:40:59
LastEditors: zmf96
LastEditTime: 2022-03-29 18:15:20
FilePath: /core/core/tasks.py
Description: 
"""

import asyncio
import base64
import inspect
import json

import config
from celery import Celery
from requests.packages import urllib3

from globals import PLUGIN_OBJECT

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


app = Celery("hotpot_tasks")

app.config_from_object(config)


@app.task
def ping():
    return {"ping": "pong"}


async def run_async_method(func):
    if inspect.iscoroutinefunction(func):
        await func()
    else:
        func()


async def _worker_entry(tool_type, **kwargs):
    try:
        Plugin = PLUGIN_OBJECT[tool_type]
        plugin_obj = Plugin(**kwargs)
        if plugin_obj.vaildator_kwargs(kwargs):
            await run_async_method(plugin_obj.init)
            await run_async_method(plugin_obj.run)
            await run_async_method(plugin_obj.close)
        return plugin_obj.results
    except Exception as e:
        print(e)
        print(Plugin)
        return None


@app.task
def worker_entry(tool_type, **kwargs):
    print(kwargs)
    loop = asyncio.get_event_loop()
    result_list = loop.run_until_complete(_worker_entry(tool_type, **kwargs))
    return result_list


if __name__ == "__main__":
    app.start()
