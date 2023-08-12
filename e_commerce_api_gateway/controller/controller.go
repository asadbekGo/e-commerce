package controller

import (
	pb "e_commerce/e_commerce_protos/product_service"
)

type Controller struct {
	productCon pb.ProductServiceClient
}

func NewController(productClient pb.ProductServiceClient) *Controller {
	return &Controller{
		productCon: productClient,
	}
}
