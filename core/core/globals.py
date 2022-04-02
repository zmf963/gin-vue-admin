#!/usr/bin/env python
# coding=utf-8

"""
Version: 0.1
Autor: zmf96
Email: zmf96@qq.com
Date: 2022-03-22 15:30:56
LastEditors: zmf96
LastEditTime: 2022-03-29 18:13:22
FilePath: /core/core/globals.py
Description: 
"""
import importlib
import os
import tarfile
import shutil
from queue import PriorityQueue, Queue

from model.model import Domain, EmailInfo, PathInfo, PortInfo
from config import MODULE_LIST

# 任务队列
TASKQUEUE = PriorityQueue(maxsize=200)

# 任务执行结果队列
RESULTQUEUE = Queue(maxsize=200)


MODEL_OBJECT = {
    "domain": Domain,
    "port_info": PortInfo,
    "path_info": PathInfo,
    "email_info": EmailInfo,
}


PLUGIN_OBJECT = {}


def load_plugin():
    """
    加载插件

    支持两种插件形式：
    1. base目录下文件中的PluginClass类型
    2. pip包lib.py中的PluginClass类型
    """

    plugin_dir = os.path.join(os.path.dirname(os.path.abspath(__file__)), "plugin")

    load_plugin_from_base(os.path.join(plugin_dir, "base"))

    load_plugin_from_config()

def load_plugin_from_config():
    for plugin_name in MODULE_LIST:
        try:
            module = importlib.import_module(plugin_name)
            PLUGIN_OBJECT[module.PluginClass.plugin_name] = module.PluginClass
        except Exception as e:
            print(e)

def load_plugin_from_base(base_dir: str) -> None:
    for file in os.listdir(base_dir):
        if file.endswith(".py") and not file.startswith("_"):
            try:
                module = importlib.import_module("plugin.base." + file[:-3])
                PLUGIN_OBJECT[module.PluginClass.plugin_name] = module.PluginClass
            except Exception as e:
                print(e)


def load_plugin_from_module_dir(plugin_dir) -> None:
    for file in os.listdir(os.path.join(plugin_dir, "module")):
        print(file)
        module = importlib.import_module("plugin.module." + file)
        print(module)
        print(dir(module))
        PLUGIN_OBJECT[module.PluginClass.plugin_name] = module.PluginClass


def load_plugin_from_pip(plugin_dir):
    def _untar(fname, dirs):
        t = tarfile.open(fname)
        t.extractall(path=dirs)
        t.close()

    for file in os.listdir(os.path.join(plugin_dir, "pip_tar_gz")):
        if file.endswith(".tar.gz"):
            tar_name = os.path.join(plugin_dir, "pip_tar_gz", file)
            module_name = file.split("-")[0]
            module_path = os.path.join(plugin_dir, "module", module_name)
            print(tar_name)
            print(module_path)
            if os.path.exists(module_path):
                shutil.rmtree(module_path)
            _untar(tar_name, os.path.join(plugin_dir, "pip_tar_gz"))
            shutil.move(os.path.join(tar_name[:-7], module_name), module_path)



load_plugin()

print(PLUGIN_OBJECT)
if __name__ == "__main__":
    for k,v in PLUGIN_OBJECT.items():
        print(k)
        print(v.usage)