FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates && \
    apk add tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone


ADD ./wechatServer /go/bin/wechatServer
ADD ./config /go/bin/config

CMD ["/go/bin/wechatServer"]
EXPOSE 10086