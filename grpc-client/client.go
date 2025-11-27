package main

import (
	"context"
	"grpc-server/greet"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// callRpc1()

	callRpc2()
}

// go-zero方式
func callRpc1() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := greet.NewGreetClient(conn.Conn())

	resp, err := client.Ping(context.Background(), &greet.Request{Ping: "ping1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(resp)
}

// 原生方式
func callRpc2() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yml", &clientConf)
	conn, err := grpc.NewClient(clientConf.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := greet.NewGreetClient(conn)
	resp, err := client.Ping(context.Background(), &greet.Request{Ping: "ping2"})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(resp)
}
