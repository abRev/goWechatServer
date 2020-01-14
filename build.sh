#!/bin/bash

param_length="$#"
if [ $param_length == 0 ]
then
  tag="test"
else
  tag="$1"
fi

docker rmi wechat_server:$tag

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wechatServer .;
docker build -t wechat_server:$tag .
echo "build success."