<!--
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-18 10:31:39
 * @LastEditors: zmf96
 * @LastEditTime: 2022-03-31 02:16:44
 * @FilePath: /quick-start_v2.md
 * @Description: 
-->



# Quick Start
- [1. 项目结构](#1-项目结构)
- [2. 环境搭建](#2-环境搭建)
- [3. 添加一个插件](#3-添加一个插件)
  - [3.4. 测试效果](#34-测试效果)
## 1. 项目结构

```
.
├── core    // 扫描Master及Worker节点
├── data  // 数据持久化目录
│   ├── mongo
│   ├── mysql
│   ├── rabbitmq
│   └── redis
├── docker-compose-local.yaml
├── docker-compose.yaml  // 线上使用的docker-compose.yml
├── server  // 后端代码
└── web  // 前端代码

```

## 2. 环境搭建

基本环境:
    - Ubuntu 20.04.3 LTS
    - Docker >= 20.10.12
    - Golang >= 1.17
    - Python >= 3.8.10
    - Nodejs >= 16.13.2

```bash 

git clone https://github.com/JDArmy/CollaborativePlatform.git
cd CollaborativePlatform

修改配置文件中的默认密码及ip地址 `server/config.docker.yaml` `core/core/common/local_config.py`

docker-compose up -d

# 开发模式下
docker-compose -f docker-compose-local.yml up -d
cd web
yarn 
yarn run serve
cd ../server
go mod tidy 
go run main.go
```

# 启动Master及Worker
```bash
cd core
pip3 install poetry 
poetry install
poetry shell
cd core
python task_monitoring.py

# 在worker节点上启动任务执行进程
celery  -A tasks worker -l info
```

## 3. 添加一个插件

目前支持两种格式的插件

- 单文件的基础插件
- 符合格式的python包




cp_common

pip install cp_common


poetry shell
poetry install
poetry update
poetry publish -u 用户名 -p 密码 -r 源码仓库地址


config.py 
```py
MODULE_LIST = ["subdomain","hotfinger"]
```
`python -m hotfinger -h`
`python -m subdomain -d jd.com -f test.txt`

### 3.4. 测试效果

待补充
