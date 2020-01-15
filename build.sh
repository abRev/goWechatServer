#!/bin/bash

param_length="$#"
if [ $param_length == 0 ]
then
  tag="test"
else
  tag="$1"
fi


CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dest/wechatServer .;
docker build -t abang/go_wechat_server:$tag .
# docker push abang/go_wechat_server:$tag
echo "build success."