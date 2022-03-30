#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-23 20:11:43
LastEditors: zmf96
LastEditTime: 2022-03-24 14:51:09
FilePath: /core/core/plugin/base/gettitle.py
Description: 
"""
import requests
from bs4 import BeautifulSoup

from cp_common.base import Plugin, Requests
from cp_common.utils import get_host_port_from_url


class PluginClass(Plugin, Requests):
    usage = [
        {
            "name": "url_list",
            "type": "List[str]",
            "usage": '域名列表，eg: ["http://baidu.com","http://bing.com"]',
        }
    ]
    plugin_name = "gettitle"
    author = "example"
    version = "0.1.0"

    def __init__(self, **kwargs):
        print(kwargs)
        super(PluginClass, self).__init__(**kwargs)
        super(Plugin, self).__init__()

    def get_title(self, url):
        """
        获取title
        """
        try:
            headers = {
                "sec-ch-ua": "Microsoft Edge",
                "sec-ch-ua-platform": "macOS",
                "Upgrade-Insecure-Requests": "1",
                "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44",
                "Accept": "text/html,application/xhtml+xml,application/xml;",
                "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
                "Connection": "close",
            }
            response = self.session.get(
                url, headers=headers, timeout=7, verify=False
            )
            soup = BeautifulSoup(response.text, "lxml")
            title = soup.find("title")
            if title and len(title.text) > 0:
                title = title.text
            else:
                title = response.text[:100]

            status_code = response.status_code
        except (
            requests.exceptions.ReadTimeout,
            requests.exceptions.ConnectTimeout,
        ) as e:
            title = "Timeout"
            status_code = 408
        except Exception as e:
            title = str(e)
            status_code = 400

        return title, str(status_code)

    def run(self, **args):
        url_list = self.kwargs.get("url_list")
        for url in url_list:
            host, port = get_host_port_from_url(url)
            hostinfo = host + ":" + port
            title, status_code = self.get_title(url)
            self.results.append(
                {
                    "port_info": {
                        "url": url,
                        "hostinfo": hostinfo,
                        "title": title,
                        "status": status_code,
                        "host": host,
                        "port": port,
                    }
                }
            )


if __name__ == "__main__":
    import fire

    # url_list = ["https://wwww.jd.com:8080", "http://bing.com"]
    # plugin = GetTitle(url_list=url_list)
    # plugin.run()
    # print(plugin.results)
    fire.Fire(PluginClass)
