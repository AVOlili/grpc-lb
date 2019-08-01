package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-lb/balancer"
	"grpc-lb/examples/proto"
	registry "grpc-lb/registry/zookeeper"
	"log"
	"time"
)

func main() {
	registry.RegisterResolver("zk", []string{"144.202.111.210:2189"}, "test", "v1.0")
	c, err := grpc.Dial("zk:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()
	client := proto.NewTestClient(c)

	for i := 0; i < 5000; i++ {
		resp, err := client.Say(context.Background(), &proto.SayReq{Content: "round robin"})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		log.Printf(resp.Content)
	}
}
