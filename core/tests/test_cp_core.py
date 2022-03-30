#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-22 14:50:56
LastEditors: zmf96
LastEditTime: 2022-03-29 18:25:13
FilePath: /core/tests/test_cp_core.py
Description: 
'''
from core import __version__
from core import tasks
from cp_common.log import logger

def test_version():
    assert __version__ == '0.1.0'


def test_task():
    rets = []
    for _ in range(10):
        ret = tasks.ping.delay()
        rets.append(ret)
        ret = tasks.worker_entry.delay("PING",host_list=["localhost"])
        rets.append(ret)
    for ret in rets:
        logger.info(ret.get())