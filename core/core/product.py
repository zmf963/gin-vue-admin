#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-14 10:37:00
LastEditors: zmf96
LastEditTime: 2022-03-07 03:00:45
FilePath: /core/core/product.py
Description: 
'''

import pika

from common.config import broker_url

connection = pika.BlockingConnection(pika.URLParameters(broker_url[2:]))
channel = connection.channel()
channel.queue_declare(queue='server:default', durable=False)
channel.basic_publish(exchange='',
                      routing_key='server:default',
                      body="ping")

print(" [x] Sent 'ping'")
connection.close()
