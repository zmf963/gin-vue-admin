#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:40:59
LastEditors: zmf96
LastEditTime: 2022-02-09 12:04:03
FilePath: /core/core/tasks.py
Description: 
'''

import base64
import json

import requests
from log import logger
from celery import Celery
from requests.packages import urllib3
import config
from config import SAVE_BASE_URL, SAVE_Headers
from plugins.gettitle.gettitle import get_title

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


app = Celery('hotpot_tasks')

app.config_from_object(config)


def save_data(path, data):
    data = json.dumps(data)
    # response = requests.request(
    #     "POST", SAVE_BASE_URL+path, headers=SAVE_Headers, data=data, verify=False)
    logger.info(data)


@app.task
def ping():
    return {"ping": "pong"}


@app.task
def nmap_scan(target):
    return target


@app.task
def icpsearch(content):
    return content

@app.task
def domainsearch(content):
    return content

@app.task
def screenshot(content):
    return content

@app.task
def appsearch(content):
    return content

@app.task
def gettitle(insert_id, url):
    title, header, body = get_title(url)
    print(insert_id)
    print(title)
    # TODO  保存结果到数据库中
    save_port_info_url = SAVE_BASE_URL + "save_port_info"
    body = str(base64.b64encode(body.encode()))
    data = {
        "_id": insert_id,
        "title": title,
        "header": str(header),
        "body": body
    }
    return data


@app.task
def cdncheck(content):
    return content

@app.task
def vscan(content):
    return content




if __name__ == '__main__':
    app.start()
