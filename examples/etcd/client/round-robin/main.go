package main

import (
	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-lb/balancer"
	"grpc-lb/examples/proto"
	registry "grpc-lb/registry/etcd"
	"log"
	"time"
)

func main() {
	etcdConfg := etcd.Config{
		Endpoints: []string{"http://144.202.111.210:2379"},
	}
	registry.RegisterResolver("etcd", etcdConfg, "test", "v1.0")

	c, err := grpc.Dial("etcd:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
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
