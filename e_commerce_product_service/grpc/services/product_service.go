package services

import (
	"context"
	"database/sql"
	"e_commerce/e_commerce_product_service/storage"
	pb "e_commerce/e_commerce_protos/product_service"
	"fmt"
)

type productService struct {
	sqlDB *sql.DB
	pb.UnimplementedProductServiceServer
}

func NewProductService(sqlDB *sql.DB) *productService {
	return &productService{
		sqlDB: sqlDB,
	}
}

func (p *productService) Create(ctx context.Context, req *pb.CreateProduct) (*pb.Product, error) {

	id, err := storage.Create(p.sqlDB, req)
	if err != nil {
		fmt.Println("Product service create: ", err)
		return nil, err
	}

	product, err := storage.GetByID(p.sqlDB, &pb.ProductPrimaryKey{Id: id})
	if err != nil {
		fmt.Println("Product service get by id: ", err)
		return nil, err
	}

	return product, nil
}

func (p *productService) GetByID(ctx context.Context, req *pb.ProductPrimaryKey) (*pb.Product, error) {

	product, err := storage.GetByID(p.sqlDB, req)
	if err != nil {
		fmt.Println("Product service get by id: ", err)
		return nil, err
	}

	return product, nil
}
func (p *productService) GetList(context.Context, *pb.GetListProductRequest) (*pb.GetListProductResponse, error) {

	return nil, nil
}
func (p *productService) Update(context.Context, *pb.Product) (*pb.Product, error) {

	return nil, nil
}
func (p *productService) Delete(context.Context, *pb.ProductPrimaryKey) (*pb.Empty, error) {

	return nil, nil
}
