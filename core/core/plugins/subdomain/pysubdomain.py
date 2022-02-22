#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-22 04:03:02
LastEditors: zmf96
LastEditTime: 2022-02-22 04:20:19
FilePath: /core/core/plugins/subdomain/pysubdomain.py
Description: 
'''

from subdomain.subdomain import SubDomain

def run_pysubdomain(domain):
    sd = SubDomain(({'deep': 1, 'domain': domain,
                   'dictname': 'test.txt'}))
    sd.run()
    return sd.results

if __name__ == '__main__':
    domain = 'baidu.com'
    print(run_pysubdomain(domain))