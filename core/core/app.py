#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 18:15:35
LastEditors: zmf96
LastEditTime: 2022-02-09 14:44:47
FilePath: /core/core/app.py
Description: 
'''

from locale import resetlocale
import tasks

if __name__ == '__main__':
    rets = []
    for i in range(1000):
        # print(tasks.ping.delay().get())
        # res = tasks.gettitle.delay(i, "https://baidu.com")
        # rets.append(res)
        res = tasks.icpsearch.delay("https://baidu.com")
        rets.append(res)

    for ret in rets:
        ret.get()