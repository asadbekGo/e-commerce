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

	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

func Load() Config {

	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	cfg.ServiceHost = cast.ToString(getDefaultValueOrReturnEnv("SERVICE_HOST", "localhost"))
	cfg.ServicePort = cast.ToString(getDefaultValueOrReturnEnv("SERVICE_PORT", ":9101"))

	cfg.PostgresHost = cast.ToString(getDefaultValueOrReturnEnv("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getDefaultValueOrReturnEnv("POSTGRES_PORT", "5432"))
	cfg.PostgresDatabase = cast.ToString(getDefaultValueOrReturnEnv("POSTGRES_DATABASE", "product_service"))
	cfg.PostgresUser = cast.ToString(getDefaultValueOrReturnEnv("POSTGRES_USER", "asadbek"))
	cfg.PostgresPassword = cast.ToString(getDefaultValueOrReturnEnv("POSTGRES_PASSWORD", "3066586"))

	return cfg
}

func getDefaultValueOrReturnEnv(key string, dafaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return dafaultValue
}
