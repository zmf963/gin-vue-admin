#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:43:49
LastEditors: zmf96
LastEditTime: 2022-02-08 17:47:52
FilePath: /core/core/log.py
Description: 
'''

import sys
import pathlib

from loguru import logger

# 路径设置
relative_directory = pathlib.Path(__file__).parent.parent  # 代码相对路径
log_path = relative_directory.joinpath('logs/log.log')  # 日志文件路径

# 日志配置
# 终端日志输出格式
stdout_fmt = '<cyan>{time:HH:mm:ss,SSS}</cyan> ' \
             '[<level>{level: <5}</level>] ' \
             '<blue>{module}</blue>:<cyan>{line}</cyan> - ' \
             '<level>{message}</level>'
# 日志文件记录格式
logfile_fmt = '<light-green>{time:YYYY-MM-DD HH:mm:ss,SSS}</light-green> ' \
              '[<level>{level: <5}</level>] ' \
              '<cyan>{process.name}({process.id})</cyan>:' \
              '<cyan>{thread.name: <18}({thread.id: <5})</cyan> | ' \
              '<blue>{module}</blue>.<blue>{function}</blue>:' \
              '<blue>{line}</blue> - <level>{message}</level>'

logger.remove()
logger.level(name='TRACE', color='<cyan><bold>', icon='✏️')
logger.level(name='DEBUG', color='<blue><bold>', icon='🐞 ')
logger.level(name='INFOR', no=20, color='<green><bold>', icon='ℹ️')
logger.level(name='QUITE', no=25, color='<green><bold>', icon='🤫 ')
logger.level(name='ALERT', no=30, color='<yellow><bold>', icon='⚠️')
logger.level(name='ERROR', color='<red><bold>', icon='❌️')
logger.level(name='FATAL', no=50, color='<RED><bold>', icon='☠️')

# 命令终端日志级别默认为INFOR
logger.add(sys.stderr, level='INFOR', format=stdout_fmt, enqueue=True)
# 日志文件默认为级别为DEBUG
logger.add(log_path, level='DEBUG', format=logfile_fmt, enqueue=True, encoding='utf-8')
