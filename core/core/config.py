#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 12:27:29
LastEditors: zmf96
LastEditTime: 2022-02-08 15:37:01
FilePath: /core/core/config.py
Description: 配置文件模板,实际配置文件local_config.py请参考此文件
'''


def hash_password(salt, password):
    salted = password + salt
    return hashlib.sha512(salted.encode("utf-8")).hexdigest()


users = {
    "zvdw3": hash_password("this is tmp project 242", "38zhVSzd"),
    "b4Vdd": hash_password("this is tmp project 242", "485DVzd3"),
    "ff5Zf": hash_password("this is tmp project 242", "485DVzd2")
}

FOFA_EMAIL = os.getenv("FOFA_EMAIL") if os.getenv(
    "FOFA_EMAIL") else "johnbrucel11t@gmail.com"
FOFA_KEYS = os.getenv("FOFA_KEYS") if os.getenv(
    "FOFA_KEYS") else "2832ba0411deb72467da4e17a0706d65"

apikeys = ["hoted2#@cVdAs"]

MOTOR_URI = "mongodb://hotdast:o43iXVVAF4v2@localhost:27017/hotdast?authSource=admin",


BROKER_URL = "pyamqp://hotpot:He23zvw3d3D@localhost:5672/"
CELERY_RESULT_BACKEND = MOTOR_URI
CELERY_TASK_SERIALIZER = 'json'
task_serializer = 'json'
CELERY_TASK_RESULT_EXPIRES = 60 * 60 * 24
CELERY_ACCEPT_CONTENT = ["json"]
CELERY_DEFAULT_QUEUE = "default"
CELERY_TIMEZONE = 'Asia/Shanghai'
CELERY_QUEUES = {
    "default": {  # 这是上面指定的默认队列
        "exchange": "default",
        "exchange_type": "direct",
        "routing_key": "default"
    }
}


SAVE_BASE_URL = "https://114.67.109.134:9523/"
SAVE_Headers = {
    'Authorization': 'Basic enZkdzM6Mzh6aFZTemQ=',
    'Content-Type': 'application/json'
}

try:
    from local_config import *
except Exception as e:
    print(e)