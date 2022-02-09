#!/bin/bash
###
 # @Version: 0.1
 # @Autor: zmf96
 # @Email: zmf96@qq.com
 # @Date: 2022-02-08 17:40:03
 # @LastEditors: zmf96
 # @LastEditTime: 2022-02-09 14:23:08
 # @FilePath: /core/runflower.sh
 # @Description: 
### 
# celery --broker="redis://default:xsC4sdf43aX@localhost:6379/0" --result-backend="redis://default:xsC4sdf43aX@localhost:6379/0" flower --basic_auth=test:celeryBflower

celery   --broker="pyamqp://hotpot:sCviEv334Vds@localhost:5672/" --result-backend="mongodb://online:43invVDwfsd4@localhost:27017/online?authSource=admin" flower --basic_auth=test:celeryBflower