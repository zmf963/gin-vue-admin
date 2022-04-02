#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-23 10:32:32
LastEditors: zmf96
LastEditTime: 2022-03-24 14:45:21
FilePath: /core/core/plugin/base/screenshot.py
Description:
"""

from pyppeteer import launch
import asyncio
import sys
from typing import List
from cp_common.base import Plugin
from cp_common.utils import get_host_port_from_url
from cp_common.req import Requests


class PluginClass(Plugin, Requests):
    plugin_name = "screenshot"
    author = "example"
    version = "0.1.0"
    usage = [
        {
            "name": "url_list",
            "type": "List[str]",
            "usage": 'Url列表，eg: ["http://baidu.com","http://bing.com"]',
        }
    ]

    def __init__(self, **kwargs) -> None:
        super(PluginClass, self).__init__(**kwargs)
        super(Plugin, self).__init__()

        self.browser = None

    def verify_depend(self):
        try:
            browser = launch(
                headless=True, ignoreHTTPSErrors=True, args=["--no-sandbox"]
            )
            browser.close()
            return True
        except Exception as e:
            self.logger.warning(e)
            return False

    async def init(self):
        self.logger.info("init")
        self.browser = await launch(
            headless=True,
            ignoreHTTPSErrors=True,
            args=["--no-sandbox", "--proxy-server={}".format(self.get_proxy())],
        )
        self.logger.info("init over")

    async def close(self):
        try:
            await self.browser.close()
        except Exception as e:
            self.logger.warning(e)

    def vaild_args(self):
        try:
            url_list = self.kwargs.get("url_list")
            for url in url_list:
                if not url.startswith("http://") and not url.startswith("https://"):
                    return False
            return True
        except Exception as e:
            self.logger.warning(e)
            return False

    async def take_screenshot(self, url):
        url = f"http://{url}" if not url.startswith("http") else url
        url = url.replace("www.", "")
        self.logger.info(f"Attempting to take a screenshot of: {url}")
        #
        context = await self.browser.createIncognitoBrowserContext()
        page = await context.newPage()
        await page.setViewport({"width": 1366, "height": 768})
        png_base64 = ""
        title = ""
        try:
            page.setDefaultNavigationTimeout(15000)
            await page.setUserAgent(self.get_user_agent())
            await page.goto(url)
            png_base64 = await page.screenshot({"encoding": "base64"})
            png_base64 = "data:image/png;base64," + png_base64
            title = await page.title()
        except Exception as e:
            print(f"An exception has occurred attempting to screenshot: {url}: {e}")
        finally:
            await page.close()
            await context.close()
        return url, title, png_base64

    async def run(self, **args):
        url_list = self.kwargs.get("url_list")
        for url in url_list:
            host, port = get_host_port_from_url(url)
            url, title, data = await self.take_screenshot(url)
            print(f" {url} {title} {len(data)}")
            self.results.append(
                {
                    "port_info": {
                        "hostinfo": url,
                        "title": title,
                        "host": host,
                        "port": port,
                        "screenshot": data,
                    }
                }
            )


async def main():
    url_list = sys.argv[1:]
    for _ in range(3):
        screen = PluginClass(url_list=url_list)
        if inspect.iscoroutinefunction(screen.init):
            await screen.init()
        else:
            screen.init()
        if inspect.iscoroutinefunction(screen.run):
            await screen.run()
        else:
            screen.run()
        print(len(screen.results))
        if inspect.iscoroutinefunction(screen.close):
            await screen.close()
        else:
            screen.close()


if __name__ == "__main__":
    import asyncio
    import inspect

    asyncio.run(main())
