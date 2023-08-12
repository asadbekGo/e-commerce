package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceHost string
	ServicePort string

	ProductServiceHost string
	ProductServicePort string
}

func Load() Config {

	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	cfg.ServiceHost = cast.ToString(getDefaultValueOrReturnEnv("SERVICE_HOST", "localhost"))
	cfg.ServicePort = cast.ToString(getDefaultValueOrReturnEnv("SERVICE_PORT", ":8080"))

	cfg.ProductServiceHost = cast.ToString(getDefaultValueOrReturnEnv("PRODUCT_SERVICE_HOST", "localhost"))
	cfg.ProductServicePort = cast.ToString(getDefaultValueOrReturnEnv("PRODUCT_SERVICE_PORT", ":9101"))

	return cfg
}

func getDefaultValueOrReturnEnv(key string, dafaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return dafaultValue
}
