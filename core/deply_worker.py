#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-09 10:27:03
LastEditors: zmf96
LastEditTime: 2022-02-09 14:36:37
FilePath: /core/deply_worker.py
Description: 部署worker节点
'''
import os
from threading import Thread
import threading
from fabric import task,Connection,SerialGroup,ThreadingGroup
from loguru import logger

hosts = [
    "192.168.31.70",
    "192.168.31.69",
    "192.168.31.168",
    "192.168.31.87",
]

def ping(c):
    logger.info("ping")
    c.run("hostname")
    c.run("ip a")
    c.run("whoami")
    c.run("uptime")
    c.run("lsb_release -a")
    c.run("uname -a")
    # c.run("cat /etc/os-release")
    # c.run("cat /etc/apt/sources.list")
    logger.info("pong")

def set_zh(c):
    c.run("sudo apt install language-pack-zh-hans -y")
    c.run("locale-gen zh_CN.UTF-8")
    c.run("update-locale LANG=zh_CN.UTF-8")
    c.run("export LANG=zh_CN.UTF-8")
    c.run("export LC_ALL=zh_CN.UTF-8")
    c.run("cat /etc/default/locale")
    c.run("date")
    c.run("sudo timedatectl set-timezone Asia/Shanghai")
    c.run("sudo timedatectl set-ntp true")
    c.run("sudo timedatectl")
    c.run("date")
    c.run("sudo timedatectl set-ntp no")

def deply_worker_run(c):
    c.run("hostname")
    c.run("date")
    c.put("dist/core-0.1.0.tar.gz","/opt/core-0.1.0.tar.gz")
    with c.cd("/opt"):
        c.run("pwd")
        c.run("tar -xzvf core-0.1.0.tar.gz")
        c.run("tree")
    with c.cd("/opt/core-0.1.0"):
        c.run("pwd")
        c.run("pip install poetry -i https://pypi.tuna.tsinghua.edu.cn/simple")
        c.put("core/local_config.py","/opt/core-0.1.0/core/local_config.py")
        c.run("poetry update")

def celery_worker_run(host):
    with Connection(host,user="root",port="22",connect_kwargs={"password":os.environ["PASSWD"]}) as c:
        # deply_worker_run(c)
        with c.cd("/opt/core-0.1.0/core"):
            c.run("pwd")
            try:
                c.run("ps auxww | awk '/celery worker/ {print $2}' | xargs kill -9")
            except Exception as e:
                logger.warning(e)
            c.run("poetry run celery  -A tasks worker -l info -Q default,low,medium,height -n default@%h")


if __name__ == "__main__":
    import threading
    threads = []
    for host in hosts:
        t = threading.Thread(target=celery_worker_run,args=(host,))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()
        
