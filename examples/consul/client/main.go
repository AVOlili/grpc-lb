package main

import (
	con_api "github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-lb/balancer"
	"grpc-lb/examples/proto"
	"grpc-lb/registry/consul"
	"log"
	"time"
)

//http://144.202.111.210:8500/v1/agent/services
func main() {
	consul.RegisterResolver("consul", &con_api.Config{Address: "http://144.202.111.210:8500"}, "test_v1.0")
	c, err := grpc.Dial("consul:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()

	client := proto.NewTestClient(c)
	for i := 0; i < 500; i++ {

		resp, err := client.Say(context.Background(), &proto.SayReq{Content: "round robin"})
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		log.Printf(resp.Content)
	}
}
