#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 18:15:35
LastEditors: zmf96
LastEditTime: 2022-03-04 03:21:39
FilePath: /core/core/app.py
Description: 
'''

import tasks
import plugins
from common.log import logger

if __name__ == '__main__':
    rets = []
    # logger.info(plugins.gettitle.gettitle.get_title("http://www.baidu.com"))
    # for i in range(1):
    #     print(tasks.ping.delay().get())
    #     res = tasks.gettitle.delay("https://baidu.com")
    #     rets.append(res)
    #     # res = tasks.icpsearch.delay("https://baidu.com")
    #     # rets.append(res)

    # for ret in rets:
    #     ret.get()
    # ret = tasks.hotfinger.delay("https://www.baidu.com").get()
    ret = tasks.fofainfo.delay("vulhub.org.cn").get()
    logger.info(ret)