package main

import (
	"fmt"
	"wechat/config"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func main() {
	if err := config.Init(""); err != nil {
        panic(err)
	}

	router := gin.Default()

	router.Any("/", hello)
	router.Run(":" + viper.GetString("port"))
}

func hello(c *gin.Context) {

	//配置微信参数
	config := &wechat.Config{
		AppID:          viper.GetString("wechat.appid"),
		AppSecret:      viper.GetString("wechat.secrect"),
		Token:         	viper.GetString("wechat.token"),
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}