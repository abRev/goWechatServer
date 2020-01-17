#!/bin/bash
server_name="go_wechat_server"

if [ $# == 0 ] 
then
    tag="test"
else
    tag="$1"
fi

docker stop $server_name && docker rm $server_name

# docker pull abang/$server_name:$tag

docker run -d --network wechat_goWechat --name $server_name -p 10086:80 abang/$server_name:$tag 

docker logs -f --tail 20 $server_name