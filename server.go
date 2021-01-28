package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "nuxt-grpc-server/common"
	"nuxt-grpc-server/services"
)

func main() {

	fmt.Println("start server")
	rpcServer := grpc.NewServer()
	services.RegisterKbInfoServiceServer(rpcServer,new(services.KbInfoService))

	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(rpcServer.Serve(listen))
}
