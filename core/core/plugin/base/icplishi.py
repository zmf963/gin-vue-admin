#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-23 20:42:54
LastEditors: zmf96
LastEditTime: 2022-03-24 17:34:35
FilePath: /core/core/plugin/base/icplishi.py
Description: 
"""
import re

from bs4 import BeautifulSoup
from cp_common.base import Plugin, Requests
from lxml import etree


class PluginClass(Plugin, Requests):
    plugin_name = "icplishi"
    author = "example"
    version = "0.1.0"
    usage = [
        {
            "name": "keyword",
            "type": "str",
            "usage": '关键字, eg: "京ICP备12345678号"',
        }
    ]

    def __init__(self, **kwargs) -> None:
        super(PluginClass, self).__init__(**kwargs)
        super(Plugin, self).__init__()
        self.BASEURL = "https://icplishi.com/"
        self.domain_list = []
        self.target = kwargs["keyword"]
        self.mode, self.target = self.check_mode()

    def get_domain(self, html):
        soup = BeautifulSoup(html, "lxml")
        text = str(soup.find_all("td"))
        text = etree.HTML(text=text)
        text = text.xpath("string(.)")
        text = re.compile(r".*\w+\.\w+").findall(text)
        self.domain_list.extend(text)

    def check_mode(self):
        if re.search(r"(http|https)\:\/\/", self.target):
            target = re.sub(r"""(http|https)://""", "", self.target)
            return "domain", target
        elif re.search(r".ICP.\d{6}.-\d{1,}", self.target):
            target = re.sub("-\d{1,}", "", self.target)  # 提取主备案号
            return "icp", target
        elif re.search(r".*\w+\.\w+", self.target):
            return "domain", self.target
        elif re.search(r".ICP.\d{6}.", self.target):
            return "icp", self.target
        else:
            self.logger.info("请检查输入是否正确")

    def get_icp(self, html):
        global number
        soup = BeautifulSoup(html, "lxml")
        text = str(
            soup.select(
                "body > div > div.container > div > div.module.mod-panel > div.bd > div:nth-child(2) > div.c-bd"
            )
        )
        text = etree.HTML(text=text)
        text = text.xpath("string(.)")
        text = re.search(r".ICP.\d{6}.", text)  # 获取备案号
        if text != None:
            number = text.group()
            return text.group()
        else:
            self.logger.warning("请检查输入是否正确")
            return ""

    def run(self):
        url = self.BASEURL + self.target
        try:
            headers = {"User-Agent": self.get_user_agent()}
            r = self.session.get(
                url=url, headers=headers, allow_redirects=True, timeout=5, verify=False
            )
            r.raise_for_status()
        except Exception as e:
            self.logger.info(e)
        if self.mode == "domain":
            try:
                icp2 = self.get_icp(r.text)
                url = self.BASEURL + icp2
                r = self.session.get(
                    url=url,
                    headers=headers,
                    allow_redirects=True,
                    timeout=5,
                    verify=False,
                )
                r.raise_for_status()
            except Exception as e:
                self.logger.info(e)
                return
            self.get_domain(r.text)
        elif self.mode == "icp":
            self.get_domain(r.text)
        self.domain_list = list(set(self.domain_list))
        for domain in self.domain_list:
            self.results.append({"domain": {"domain": domain, "source": "icplishi"}})

    @staticmethod
    def vaildator_args(kwargs):
        if "keyword" in kwargs and isinstance(kwargs["keyword"], str):
            return True
        return False


if __name__ == "__main__":
    keyword = "https://www.jd.com"
    plugin = PluginClass(keyword=keyword)
    plugin.run()
    print(plugin.results)
