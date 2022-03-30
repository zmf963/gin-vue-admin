#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-22 17:41:24
LastEditors: zmf96
LastEditTime: 2022-03-25 17:41:20
FilePath: /core/core/plugin/base/example.py
Description: 
"""

from cp_common.base import Plugin, Requests


class PluginClass(Plugin, Requests):
    usage = [
        {
            "name": "url_list",
            "type": "List[str]",
            "usage": 'url列表，eg: ["http://baidu.com","http://bing.com"]',
        }
    ]
    plugin_name = "example"
    author = "example"
    version = "0.1.0"

    def __init__(self, **kwargs):
        super().__init__(kwargs)

    def run(self):
        self.logger.info(self.kwargs)
        self.results.append({})
