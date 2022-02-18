#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 12:27:34
LastEditors: zmf96
LastEditTime: 2022-02-15 17:37:12
FilePath: /core/local_config.py
Description: 本地配置文件
'''

import os
import hashlib

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

MOTOR_URI = "mongodb://online:43invVDwfsd4@192.168.31.24:27017/online?authSource=admin"


# broker_url = "redis://default:xsC4sdf43aX@192.168.31.24:6379/0"
broker_url = "pyamqp://hotpot:sCviEv334Vds@192.168.31.24:5672/"
# result_backend = broker_url
result_backend = MOTOR_URI
task_serializer = 'json'
result_expires = 60 * 60 * 24
accept_content = ["json"]
task_default_queue = "default"
timezone = 'Asia/Shanghai'


SAVE_BASE_URL = "https://114.67.109.134:9523/"
SAVE_Headers = {
    'Authorization': 'Basic enZkdzM6Mzh6aFZTemQ=',
    'Content-Type': 'application/json'
}