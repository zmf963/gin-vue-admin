#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-23 03:35:06
LastEditors: zmf96
LastEditTime: 2022-03-07 10:42:29
FilePath: /core/core/consum_data_done.py
Description: 
'''

from core.model import EmailInfo
from common.log import logger
from model import Task, PathInfo, Domain, PortInfo
from bson.objectid import ObjectId

def consum_gettile(data, task_obj):
    try:
        tmp = data.get("url").split("://")[1].split("/")
        data["hostinfo"] = tmp[0]
        data["path"] = "/"+"/".join(tmp[1:])
        data["target_id"] = task_obj.target_id
        data["response_size"] = len(data.get("body"))
        if PathInfo.objects(url=data.get("url")).count() > 0:
            PathInfo.objects(url=data.get("url")).first().Update(**data)
        else:
            PathInfo(**data).Save()
        if data["path"] == "/":
            data.pop("path")
            data.pop("response_size")
            url = data.pop("url")
            tmp = data.get("hostinfo").split(":")
            data["host"] = tmp[0]
            if len(tmp) > 1:
                port = tmp[1]
            elif url.startswith("https://"):
                port = "443"
            else:
                port = "80"
            data["port"] = port
            if PortInfo.objects(hostinfo=data.get("hostinfo")).count() > 0:
                PortInfo.objects(hostinfo=data.get(
                    "hostinfo")).first().Update(**data)
            else:
                PortInfo(**data).Save()
    except Exception as e:
        logger.warning(task_obj)
        logger.warning(e)


def consum_cdncheck(data, task_obj):
    for item in data:
        logger.info(item)
        item["target_id"] = task_obj.target_id
        cdn = 0
        if item.get("cdn") != '':
            cname = item.get("cdn")
            cdn = 1
        if Domain.objects(domain=item.get("target")).count() > 0:
            do = Domain.objects(domain=item.get("target")).first()
            if item.get("ip") not in do.ips:
                do.ips.append(item.get("ip"))
            do.cdn = cdn
            do.cname = cname
            do.Save()
        else:
            Domain(domain=item.get("target"), cdn=cdn,
                   cname=cname, ips=[item.get("ip")]).Save()


def consum_beian2domain(data, task_obj):
    for item in data:
        logger.info(item)
        tmp = {"domain": item[1],
               "target_id": task_obj.target_id, "tags": [item[0]]}
        if Domain.objects(domain=item[1]).count() > 0:
            do = Domain.objects(domain=item[1]).first()
            if item[0] not in do.tags:
                do.tags.append(item[0])
                do.Save()
        else:
            Domain(**tmp).Save()


def consum_pysubdomain(data, task_obj):
    for k, v in data.items():
        logger.info(k, v)
        print(task_obj.target_id)
        tmp = {"domain": k, "target_id": ObjectId(task_obj.target_id), "ips": v}
        if Domain.objects(domain=k).count() > 0:
            # do = Domain.objects(domain=k).first()
            pass
        else:
            Domain(**tmp).Save()


def consum_hotfinger(data, task_obj):
    try:
        tmp = data.get("domain").split("://")[1].split("/")
        data["hostinfo"] = tmp[0]
        data["target_id"] = task_obj.target_id
        data["products"] = data.pop("fingers")
        tmp = data.get("hostinfo").split(":")
        data["host"] = tmp[0]
        if len(tmp) > 1:
            port = tmp[1]
        elif data.get("domain").startswith("https://"):
            port = "443"
        else:
            port = "80"
        data["port"] = port
        print(data)
        if PortInfo.objects(hostinfo=data.get("hostinfo")).count() > 0:
            PortInfo.objects(hostinfo=data.get(
                "hostinfo")).first().Update(**data)
        else:
            PortInfo(**data).Save()
    except Exception as e:
        logger.warning(task_obj)
        logger.warning(e)


def consum_fofainfo(data, task_obj):
    if data.get("error") == False:
        for result in data.get("results"):
            if Domain.objects(domain=result[1]).count() > 0:
                do = Domain.objects(domain=result[1]).first()
                if result[3] not in do.ips:
                    do.ips.append(result[3])
                do.addr = " ".join(result[5:8])
                do.cname = result[-1]
                do.source = "fofa"
                if result[-2] not in do.tags:
                    do.tags.append(result[-2])
                do.Save()
            else:
                Domain(domain=result[1], addr=" ".join(result[5:8]),
                       cname=result[-1], ips=[result[3]], tags=[result[-2]],
                       source="fofa", target_id=task_obj.target_id
                       ).Save()

            if ":" in result[1]:
                hostinfo = result[1]
            else:
                hostinfo = result[1] + result[4]
            if PortInfo.objects(hostinfo=hostinfo).count() > 0:
                pi = PortInfo.objects(hostinfo=hostinfo).first()
                pi.title = result[0]
                pi.port = result[4]
                if result[2] == "":
                    pi.host = result[3]
                else:
                    pi.host = result[2]
                if result[-3] not in pi.products:
                    pi.products.append(result[-3])
                pi.Save()
            else:
                PortInfo(hostinfo=hostinfo, title=result[0], port=result[4],
                         host=result[3], products=[result[-3]],
                         target_id=task_obj.target_id).Save()
            PortInfo.objects


def consum_emailall(data, task_obj):
    for email in data:
        logger.info(email)
        if EmailInfo.objects(email=email).count() == 0:
            EmailInfo(email=email, target_id=task_obj.target_id).Save()


def consum_data_done(tool_type, data, task_obj):
    if tool_type == "gettitle":
        tmp_data = data[0]
        consum_gettile(tmp_data, task_obj)
    elif tool_type == "cdncheck":
        consum_cdncheck(data, task_obj)
    elif tool_type == "beian2domain":
        consum_beian2domain(data, task_obj)
    elif tool_type == "pysubdomain":
        consum_pysubdomain(data, task_obj)
    elif tool_type == "hotfinger":
        consum_hotfinger(data, task_obj)
    elif tool_type == "fofainfo":
        consum_fofainfo(data, task_obj)
    elif tool_type == "emailall":
        consum_emailall(data, task_obj)
