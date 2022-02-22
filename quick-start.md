<!--
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-18 10:31:39
 * @LastEditors: zmf96
 * @LastEditTime: 2022-02-22 07:21:57
 * @FilePath: /quick-start.md
 * @Description: 
-->



# Quick Start
- [1. 项目结构](#1-项目结构)
- [2. 环境搭建](#2-环境搭建)
- [3. 添加一个插件](#3-添加一个插件)
  - [3.1. 插件主体部分](#31-插件主体部分)
    - [3.1.1. 创建目录及文件](#311-创建目录及文件)
    - [3.1.2. 编写插件代码](#312-编写插件代码)
  - [3.2. 注册Task](#32-注册task)
  - [3.3. 结果清洗](#33-结果清洗)
  - [3.4. 测试效果](#34-测试效果)
## 1. 项目结构

```
.
├── core    // 扫描Master及Worker节点
│   ├── core
│   ├── create_ssl_tls.sh
│   ├── deply_worker.py
│   ├── dist
│   ├── docker-compose.yml
│   ├── Dockerfile
│   ├── logs
│   ├── poetry.lock
│   ├── pyproject.toml
│   ├── quick-start.md
│   ├── README.rst
│   ├── runflower.sh
│   ├── tests
│   ├── t.json
│   └── t.py
├── create_ssl_tls.sh
├── data  // 数据持久化目录
│   ├── mongo
│   ├── mysql
│   ├── rabbitmq
│   └── redis
├── docker-compose-local.yaml
├── docker-compose.yaml  // 线上使用的docker-compose.yml
├── LICENSE
├── README.md
├── SECURITY.md
├── server  // 后端代码
│   ├── api
│   ├── config
│   ├── config.docker.yaml
│   ├── config.yaml
│   ├── core
│   ├── Dockerfile
│   ├── docs
│   ├── global
│   ├── go.mod
│   ├── go.sum
│   ├── initialize
│   ├── log
│   ├── main.go
│   ├── middleware
│   ├── model
│   ├── packfile
│   ├── plugin
│   ├── README.md
│   ├── resource
│   ├── router
│   ├── server
│   ├── service
│   ├── source
│   └── utils
└── web  // 前端代码
    ├── babel.config.js
    ├── dist
    ├── docker-compose.yml
    ├── Dockerfile
    ├── favicon.ico
    ├── index.html
    ├── limit.js
    ├── node_modules
    ├── openDocument.js
    ├── package.json
    ├── package-lock.json
    ├── README.md
    ├── src
    ├── vite.config.js
    └── yarn.lock

31 directories, 37 files
```

## 2. 环境搭建

基本环境:
    - Ubuntu 20.04.3 LTS
    - Docker >= 20.10.12
    - Golang >= 1.17
    - Python >= 3.8
    - Nodejs >= 16.13.2

```bash 

git clone https://github.com/JDArmy/CollaborativePlatform.git
cd CollaborativePlatform

修改配置文件中的默认密码!!!
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
python task_watch.py # 启动任务分发
python task_task_done.py # 启动数据清洗

# 在worker节点上启动任务执行进程
celery  -A tasks worker -l info -c 100 -Q default,low,medium,height -n default@%h 
```

## 3. 添加一个插件

插件的主要代码存放core/core/plugins目录下

下面以添加一个subdomain模块为例

### 3.1. 插件主体部分


#### 3.1.1. 创建目录及文件

```bash
mkdir subdomain && cd subdomain
touch __init__.py
touch pysubdomain.py
```

#### 3.1.2. 编写插件代码

pysubdomain.py

```py
from subdomain.subdomain import SubDomain

def run_pysubdomain(domain):
    sd = SubDomain(({'deep': 1, 'domain': domain,
                   'dictname': 'default.txt'}))
    sd.run()
    return sd.results

if __name__ == '__main__':
    domain = 'baidu.com'
    print(run_pysubdomain(domain))
```

该插件使用subdomain库, 使用`poetry add subdomain`命令添加依赖


### 3.2. 注册Task

修改`core/core/tasks.py`文件,在celery中添加一个task

```py
@app.task
def pysubdomain(domain):
    result = run_pysubdomain(domain)
    data = {
        "tool_type": "pysubdomain",
        "data": result,
        "status": "complete",
    }
    return data
```

修改`core/core/task_watch.py`文件,监听pysubdomain类型任务的下发.

在callback函数中添加如下代码

```py
elif tool == "pysubdomain":
    for host in task.get("hosts").split(","):
        res = tasks.pysubdomain.delay(host)
        logger.inof(res)
        celery_task_ids.append(res.id)
```

### 3.3. 结果清洗

修改`core/core/consum_task_done.py`文件,处理任务执行结果

添加`consum_pysubdomain(data,task_obj)`函数,整理数据并入库

修改consum_done(task)函数,添加
```
            elif res.get("tool_type") == "pysubdomain":
                consum_pysubdomain(res.get("data"),task_obj)
```

```py
def consum_pysubdomain(data,task_obj):
    for k,v in data:
        logger.info(k,v)
        tmp = {"domain":k,"target_id":task_obj.target_id,"project_id":task_obj.project_id,"ips":v}
        if Domain.objects(domain=k).count() > 0:
            # do = Domain.objects(domain=k).first()
            pass
        else:
            Domain(**tmp).Save()

def consum_done(task):
    try:
        logger.info(task)
        task_obj = Task.objects(id=task.get("_id")).first()
        task_obj.Update(status="running",celery_task_ids=task.get("celery_task_ids"))
        for task_id in task.get("celery_task_ids"):
            logger.info(task_id)
            res = app.AsyncResult(task_id).get()
            logger.info(res)
            if res.get("tool_type") == "gettitle":
                tmp_data = res["data"][0]
                consum_gettile(tmp_data,task_obj)
            elif res.get("tool_type") == "cdncheck":
                consum_cdncheck(res["data"],task_obj)
            elif res.get("tool_type") == "beian2domain":
                consum_beian2domain(res.get("data"),task_obj)
            elif res.get("tool_type") == "pysubdomain":
                consum_pysubdomain(res.get("data"),task_obj)
            else:
                pass
        task_obj.Update(status="complete")
```

### 3.4. 测试效果

重新启动`task_watch.py` `consum_task_done.py` `celery worker`

添加一个扫描任务

```
curl --location --request POST 'http://127.0.0.1:8080/api/task/createTask' \
--header 'x-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNzQ1ZDllNWMtOTA2NC00NDRlLTg2NGMtNDUxZDkyYWVmNDE3IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiQnVmZmVyVGltZSI6MCwiZXhwIjoxNjQ2MTE3MzM3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE2NDU1MTE1Mzd9.ds_UkWLbRB09xTkMEUm_kN6bpUXv7aoew3BbD8GV0CI' \
--header 'Content-Type: application/json' \
--data-raw '{
    "task_name": "pysubdoamin_001",
    "hosts": "jd.com",
    "ports": "",
    "keyword": "",
    "tools": [
        "pysubdomain"
    ],
    "status": "",
    "project_id": "",
    "target_id": ""
}'
```