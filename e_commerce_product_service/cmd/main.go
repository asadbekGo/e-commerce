package main

import (
	"fmt"
	"net"

	"e_commerce/e_commerce_product_service/config"
	"e_commerce/e_commerce_product_service/grpc"
	"e_commerce/e_commerce_product_service/storage"
)

func main() {

	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.ServiceHost+cfg.ServicePort)
	if err != nil {
		panic(err)
	}

	sqlDB, err := storage.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	server := grpc.NewGrpc(sqlDB)

	fmt.Println("Listening...", cfg.ServiceHost+cfg.ServicePort)
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
