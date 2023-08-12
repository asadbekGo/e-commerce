package storage

import (
	"database/sql"
	pb "e_commerce/e_commerce_protos/product_service"

	uuid "github.com/google/uuid"
)

func Create(sqlDb *sql.DB, req *pb.CreateProduct) (string, error) {
	query := `
		INSERT INTO product(
			id,
			name,
			image,
			price,
			discount_type,
			discount,
			barcode
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	id := uuid.New()

	_, err := sqlDb.Exec(query,
		id.String(),
		req.Name,
		req.Image,
		req.Price,
		req.DiscountType,
		req.Discount,
		req.Barcode,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func GetByID(sqlDb *sql.DB, req *pb.ProductPrimaryKey) (*pb.Product, error) {

	var product = pb.Product{}

	var query = `
		SELECT
			id,
			COALESCE(name, ''),
			COALESCE(image, ''),
			COALESCE(price, 0),
			COALESCE(discount_type, ''),
			COALESCE(discount, 0),
			COALESCE(barcode, '')
		FROM product
		WHERE id = $1
	`

	err := sqlDb.QueryRow(query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Image,
		&product.Price,
		&product.DiscountType,
		&product.Discount,
		&product.Barcode,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
