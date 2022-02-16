#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:40:59
LastEditors: zmf96
LastEditTime: 2022-02-16 17:57:29
FilePath: /core/tasks.py
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
def gettitle(url):
    title, header, body,status_code = get_title(url)
    data = {
        "tool_type": "gettitle",
        "title": title,
        "url":url,
        "header": str(header),
        "body": body,
        "status": str(status_code)
    }
    return data


@app.task
def cdncheck(content):
    return content


@app.task
def vscan(content):
    return content


@app.task
def portscan(content):
    return content


if __name__ == '__main__':
    app.start()
