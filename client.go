package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"nuxt-grpc-server/services"
)

func main() {

	conn, err := grpc.Dial(":8088",grpc.WithInsecure() )
	if err != nil {
		log.Fatal(err)
	}

	//defer conn.Close()

	kbInfoClient := services.NewKbInfoServiceClient(conn)
	kbInfoResponse, err := kbInfoClient.UpdateKbDetailList(context.Background(), &services.KbInfoRequest{Id: []int64{120}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(kbInfoResponse)
}