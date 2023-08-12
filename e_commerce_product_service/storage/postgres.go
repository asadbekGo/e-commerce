package storage

import (
	"database/sql"

	_ "github.com/lib/pq"

	"e_commerce/e_commerce_product_service/config"
	"fmt"
)

func NewConnectionPostgres(cfg *config.Config) (*sql.DB, error) {

	config := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
	)

	sqlDB, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
