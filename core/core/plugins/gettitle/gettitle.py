#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-08 17:50:22
LastEditors: zmf96
LastEditTime: 2022-02-17 15:06:34
FilePath: /core/core/plugins/gettitle/gettitle.py
Description: 
'''

import requests
from requests.models import ReadTimeoutError
import urllib3
from bs4 import BeautifulSoup

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

try:
    from common.log import logger
except Exception as e:
    print(e)

def get_title(url):
    logger.info("get_title: %s" % url)
    try:
        headers = {
            "Connection": "keep-alive",
            "sec-ch-ua": "Microsoft Edge",
            "sec-ch-ua-platform": "macOS",
            "Upgrade-Insecure-Requests": "1",
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44",
            "Accept": "text/html,application/xhtml+xml,application/xml;",
            "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
            'Connection': 'close'
        }
        response = requests.get(url, headers=headers,
                                timeout=7, verify=False, stream=False)
        soup = BeautifulSoup(response.text, 'lxml')
        title = soup.find('title')
        if title and len(title.text) > 0:
            return title.text, response.headers, response.text, response.status_code
        return response.text[:100], response.headers, response.text, response.status_code
    except requests.exceptions.ReadTimeout as e:
        return "Timeout", "", "", 408
    except requests.exceptions.ConnectTimeout as e:
        return "Timeout", "", "", 408
    except Exception as e:
        return str(e), "", "", 400


if __name__ == "__main__":
    url = "https://wwww.jd.com:8080"
    # url = "http://39.107.234.104"
    try:
        import os,sys
        basedir = os.path.dirname(os.path.abspath(__file__))
        sys.path.append(os.path.abspath(os.path.join(basedir, '../../')))
        from common.log import logger
        logger.info("main")
    except Exception as e:
        print(e)
    print(get_title(url))