#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-17 18:52:13
LastEditors: zmf96
LastEditTime: 2022-02-17 18:52:13
FilePath: /common/utils.py
Description: 
'''
from typing import List,Set
from itertools import product


def parse_host(host: str) -> List[str]:
    return host.split(",")


def parse_port(port: str) -> List[str]:
    # TODO:解析 eg:1-65535
    port_list = port.split(",")
    port_list = list(set(port_list))
    port_list.remove('')
    return port_list

def get_url_list_from_host_port_list(host_list: List[str], port_list: List[str] = []) -> Set[str]:
    url_list = set()
    for host in host_list:
        if port_list != [] and port_list != ['']:
            for port in port_list:
                if "443" in port:
                    url_list.add("https://"+host+":"+port)
                else:
                    url_list.add("http://"+host+":"+port)
        else:
            url_list.add("https://"+host)
            url_list.add("http://"+host)

    return url_list

