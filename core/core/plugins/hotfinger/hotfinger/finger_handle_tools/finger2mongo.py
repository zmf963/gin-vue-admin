#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-03 07:03:04
LastEditors: zmf96
LastEditTime: 2022-03-03 07:36:47
FilePath: /core/core/plugins/hotfinger/hotfinger/finger_handle_tools/finger2mongo.py
Description: 
'''
import requests
import json

url = "http://localhost:8080/api/finger_info/createFingerInfo"

headers = {
  'x-token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzQ1ZDllNWMtOTA2NC00NDRlLTg2NGMtNDUxZDkyYWVmNDE3IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6MCwiZXhwIjoxNjQ2ODkyOTc3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2NDYyODcxNzd9.odY6wdG07pO_j62Vsa-OSVWHHcl0arWxR7UnKDE9nkI',
  'Content-Type': 'application/json'
}



def main():
    with open("../data/hotfinger_v1.json") as fp:
        finger_list = json.load(fp)
    for finger in finger_list:
        try:
            payload = json.dumps({
                "_id":finger["_id"]["$oid"],
                "name":finger["name"],
                "content":finger["content"],
                "tags":finger["tags"]
            })
            response = requests.request("POST", url, headers=headers, data=payload)
            print(response.text)
        except Exception as e:
            print(e)
            print(finger)
if __name__ == "__main__":
    main()