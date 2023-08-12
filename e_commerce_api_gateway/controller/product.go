package controller

import (
	"context"
	"fmt"

	pb "e_commerce/e_commerce_protos/product_service"

	"github.com/gin-gonic/gin"
)

func (con *Controller) ProductCreate(c *gin.Context) {

	var productData pb.CreateProduct

	err := c.BindJSON(&productData)
	if err != nil {
		fmt.Println("error while product create:", err)
		c.JSON(400, err)
		return
	}

	resp, err := con.productCon.Create(context.Background(), &productData)

	if err != nil {
		fmt.Println("call product service create:", err)
		return
	}

	c.JSON(200, resp)
}

func (con *Controller) GetByIDProduct(c *gin.Context) {

	resp, err := con.productCon.GetByID(context.Background(), &pb.ProductPrimaryKey{Id: c.Param("id")})

	if err != nil {
		fmt.Println("call product service get by id:", err)
		return
	}

	c.JSON(200, resp)
}
