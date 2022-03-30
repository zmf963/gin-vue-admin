#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 18:15:35
LastEditors: zmf96
LastEditTime: 2022-03-30 20:19:51
FilePath: /core/core/app.py
Description: 
'''

import tasks
from cp_common.log import logger

if __name__ == '__main__':
    rets = []
    for _ in range(10):
        ret = tasks.ping.delay()
        rets.append(ret)
        ret = tasks.worker_entry.delay("gettitle",url_list=["http://bing.com:80"])
        ret = tasks.worker_entry.delay("hotfinger",url_list=["http://bing.com:80","http://vulhub.org.cn"])
        rets.append(ret)
        
    for ret in rets:
        logger.info(ret.get())