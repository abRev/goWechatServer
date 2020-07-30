package grpc

import (
	"log"

	helloworld "wechat/pb/helloworld"
	routeguide "wechat/pb/routeguide"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var HelloClient helloworld.GreeterClient
var RouteClient routeguide.RouteGuideClient

func init() {
	host := viper.GetString("grpc.host")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	HelloClient = helloworld.NewGreeterClient(conn)
	RouteClient = routeguide.NewRouteGuideClient(conn)
}
