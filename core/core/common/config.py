#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 12:27:29
LastEditors: zmf96
LastEditTime: 2022-02-23 08:13:09
FilePath: /core/core/common/config.py
Description: 配置文件模板,实际配置文件local_config.py请参考此文件
'''
import os
import hashlib



try:
    from common.local_config import *
except Exception as e:
    print(e)
    
    