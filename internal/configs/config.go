package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	Environment string
	LoggerLevel string

	PostrgresHost    string
	PostrgresPort    int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	GrpcPort string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.PostrgresHost = cast.ToString(coalesce("POSTGRES_HOST", "localhost"))
	config.PostrgresPort = cast.ToInt(coalesce("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(coalesce("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("POSTGRES_PASSWORD", "3333"))
	config.PostgresDatabase = cast.ToString(coalesce("POSTGRES_DATABASE", "carpetwash_service"))

	config.GrpcPort = cast.ToString(coalesce("GRPC_PORT", ":8082"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}
