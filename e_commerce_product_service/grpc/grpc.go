package grpc

import (
	"database/sql"
	"e_commerce/e_commerce_product_service/grpc/services"
	pb "e_commerce/e_commerce_protos/product_service"

	"google.golang.org/grpc"
)

func NewGrpc(sqlDB *sql.DB) *grpc.Server {

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, services.NewProductService(sqlDB))

	return server
}
