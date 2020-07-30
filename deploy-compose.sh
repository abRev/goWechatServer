#!/bin/bash
# 通过compose启动服务,先编译服务，在使用docker打包镜像
docker-compose stop
docker-compose rm
docker rm postgres mongo elasticsearch redis go_wechat_server_compose
# sh build.sh compose
docker-compose up -d