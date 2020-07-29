package grpc

import (
	"log"

	"wechat/pb"

	"github.com/spf13/viper"
	"goole.golang.org/grpc"
)

var Client pb.GreeterClient

func init() {
	host := viper.GetString("grpc.host")
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	Client := pb.NewGreeterClient(conn)
}
