#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:40:59
LastEditors: zmf96
LastEditTime: 2022-03-04 03:20:39
FilePath: /core/core/tasks.py
Description: 
'''

import base64
import json

import requests
from celery import Celery
from requests.packages import urllib3

import common.config as config
from common.config import SAVE_BASE_URL, SAVE_Headers
from common.log import logger
from plugins.cdncheck.check_cdn import CheckCDN
from plugins.gettitle.gettitle import get_title
from plugins.subdomain.beian2domain import run_beian2domain
from plugins.subdomain.pysubdomain import run_pysubdomain
from plugins.hotfinger.hotfinger.hotfinger import hotfinger_init,run_hotfinger
from plugins.fofainfo.fofa import FofaApi

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
    print(url)
    title, header, body, status_code = get_title(url)
    data = {
        "tool_type": "gettitle",
        "data": [
            {
                "title": title,
                "url": url,
                "header": str(header),
                "body": body,
                "status": str(status_code)
            }
        ],
        "status": "complete"
    }
    return data


@app.task
def cdncheck(host):
    if isinstance(host, str):
        host = host.split(",")
    cc = CheckCDN(host)
    cc.run()
    result = []
    while not cc.queue.empty():
        result.append(cc.queue.get())
    data = {
        "tool_type": "cdncheck",
        "data": result,
        "status": "complete"
    }
    return data


@app.task
def beian2domain(keyword):
    result = run_beian2domain(keyword)
    data = {
        "tool_type": "beian2domain",
        "data": result,
        "status": "complete",
    }
    return data


@app.task
def pysubdomain(domain):
    result = run_pysubdomain(domain)
    data = {
        "tool_type": "pysubdomain",
        "data": result,
        "status": "complete",
    }
    return data

@app.task
def pysubdomain(domain):
    result = run_pysubdomain(domain)
    data = {
        "tool_type": "pysubdomain",
        "data": result,
        "status": "complete",
    }
    return data

@app.task
def hotfinger(domain):
    hotfinger_init()
    result = run_hotfinger(domain)
    data = {
        "tool_type": "hotfinger",
        "data": result,
        "status": "complete",
    }
    return data

@app.task
def fofainfo(keyword):
    fofaapi = FofaApi(config.FOFA_EMAIL, config.FOFA_KEYS)
    fofaapi.get_info(keyword)
    data = {
        "tool_type": "fofainfo",
        "data": fofaapi.results,
        "status": "complete",
    }
    return data

@app.task
def vscan(content):
    return content


@app.task
def portscan(content):
    return content


if __name__ == '__main__':
    app.start()
