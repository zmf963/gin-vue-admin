#!/usr/bin/env python
# coding=utf-8

'''
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-02-15 16:49:13
LastEditors: zmf96
LastEditTime: 2022-03-07 09:00:17
FilePath: /core/core/model.py
Description: 
'''

from enum import unique
from importlib.metadata import requires
from sqlite3 import InterfaceError, connect
from mongoengine import connect, DynamicDocument, IntField, StringField, ListField, DictField, DateTimeField, BooleanField, ObjectIdField
import datetime

from common.config import MOTOR_URI

connect(host=MOTOR_URI)


class Task(DynamicDocument):
    task_name = StringField(required=True, unique=True)
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

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class PathInfo(DynamicDocument):
    hostinfo = StringField(required=True)
    path = StringField(requried=True)
    url = StringField(required=True, unique=True)
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

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class Domain(DynamicDocument):
    domain = StringField(required=True)
    ips = ListField(StringField())
    hostnames = ListField(StringField())
    os = StringField()
    whois = DictField()
    alive = IntField(default=0)
    cname = StringField()
    cdn = IntField(default=0)
    cidr = StringField()
    asn = StringField()
    org = StringField()
    addr = StringField()
    isp = StringField()
    source = StringField()
    target_id = ObjectIdField()
    target_id_is_verify = BooleanField(defalut=False)
    port_ids = ListField(StringField())

    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_domain'}

    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class Target(DynamicDocument):
    target_name = StringField(required=True)
    project_ids = ListField(StringField())
    task_ids = ListField(StringField())
    domain_ids = ListField(StringField())

    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_target'}

    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class Project(DynamicDocument):
    project_name = StringField(required=True)
    project_desc = StringField()
    start_time = DateTimeField(default=datetime.datetime.now())
    end_time = DateTimeField()
    target_ids = ListField(StringField())
    task_ids = ListField(StringField())

    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_project_info'}

    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class PortInfo(DynamicDocument):
    port = StringField(required=True)
    host = StringField(required=True)
    hostinfo = StringField(required=True, unique=True)
    url = StringField()
    title = StringField()
    favicons = StringField()
    screenshot = StringField()
    products = ListField(StringField())
    protocols = ListField(StringField())
    alive = IntField()
    banner = StringField()
    status = StringField()
    domain_id = StringField()
    path_ids = ListField(StringField())

    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_port_info'}

    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()


class EmailInfo(DynamicDocument):
    email = StringField(required=True, unique=True)
    source = StringField(requried=True)
    target_id = StringField(required=True)

    tags = ListField(StringField())
    remarks = StringField()
    create_at = DateTimeField(default=datetime.datetime.now())
    update_at = DateTimeField(default=datetime.datetime.now())
    delete_at = DateTimeField()

    meta = {'collection': 'pro_email_info'}

    def Update(self, **kwargs):
        self.update_at = datetime.datetime.utcnow()
        for k, v in kwargs.items():
            setattr(self, k, v)
        self.save()

    def Save(self):
        self.update_at = datetime.datetime.utcnow()
        self.save()
