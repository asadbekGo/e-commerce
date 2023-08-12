package main

import (
	"e_commerce/e_commerce_api_gateway/config"
	"e_commerce/e_commerce_api_gateway/controller"
	pb "e_commerce/e_commerce_protos/product_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cfg := config.Load()

	conn, err := grpc.Dial(
		cfg.ProductServiceHost+cfg.ProductServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	product_service := pb.NewProductServiceClient(conn)

	r := gin.New()

	con := controller.NewController(product_service)

	r.POST("/product", con.ProductCreate)
	r.GET("/product/:id", con.GetByIDProduct)

	err = r.Run(cfg.ServiceHost + cfg.ServicePort)
	if err != nil {
		panic(err)
	}
}
