#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-15 16:49:13
LastEditors: zmf96
LastEditTime: 2022-02-16 15:18:49
FilePath: /core/model.py
Description: 
'''

from sqlite3 import connect
from mongoengine import connect, DynamicDocument,IntField, StringField, ListField,DictField,DateTimeField
import datetime

from config import MOTOR_URI

connect(host=MOTOR_URI)

class MyDynamicDocument(DynamicDocument):
    meta = {"allow_inheritance": True}
    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

class Task(DynamicDocument):
    task_name = StringField(required=True)
    hosts = StringField()
    ports = StringField()
    keyword = StringField()
    tools = ListField(StringField())
    tool_ext = DictField()
    status = StringField()
    project_id = StringField()
    target_id = StringField()
    celery_task_ids = ListField(StringField(max_length=50))
    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_task'}
    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()


class PathInfo(DynamicDocument):
    hostinfo = StringField()
    path = StringField()
    url = StringField(required=True)
    title = StringField()
    status = StringField()
    response_size = IntField()
    body = StringField()
    header = StringField()
    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_path_info'}
    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()
