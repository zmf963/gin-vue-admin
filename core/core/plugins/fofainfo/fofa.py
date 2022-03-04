#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-04 02:52:07
LastEditors: zmf96
LastEditTime: 2022-03-04 03:06:05
FilePath: /core/core/plugins/fofainfo/fofa.py
Description: 
'''
import http.client
import sys
import os
import json
import base64

class FofaApi:
    def __init__(self, api_email,api_key):
        self.api_email = api_email
        self.api_key = api_key
        self.keywords = None
        self.results = None
    def get_info(self,keywords):
        if keywords is not None:
            self.keywords = keywords
        conn = http.client.HTTPSConnection("fofa.info")
        payload = ''
        headers = {}
        _keyword_base64 = str(base64.urlsafe_b64encode(self.keywords.encode('utf-8'))).lstrip('b\'').rstrip('\'')
        # _query_url = "/api/v1/search/all?email="+self.api_email+"&key="+self.api_key+"&qbase64="+_keyword_base64+"&fields=host,title,ip,domain,port,country,province,city,country_name,header,server,protocol,banner,cert,isp,as_number,as_organization,latitude,longitude,icp,fid,cname"
        _query_url = "/api/v1/search/all?email="+self.api_email+"&key="+self.api_key+"&qbase64="+_keyword_base64+"&fields=title,host,domain,ip,port,country_name,province,city,server,protocol,cname"
        # print(_query_url)
        conn.request("POST", _query_url, payload, headers)
        res = conn.getresponse()
        data = res.read()
        self.results = json.loads(data.decode("utf-8"))

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python fofa.py <keyword>")
        sys.exit()
    keyword = sys.argv[1]
    fofaapi = FofaApi(os.getenv('FOFA_EMAIL'),os.getenv('FOFA_KEY'))
    fofaapi.get_info(keyword)
    result_json = json.dumps(
        fofaapi.results, indent=4, sort_keys=True, ensure_ascii=False)
    print(result_json)
    