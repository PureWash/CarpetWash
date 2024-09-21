package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string

	PostrgresHost    string
	PostrgresPort    int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func Load() Config {
	if err := godotenv.Load(); err != nil{
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_Port", ":8082"))

	config.PostrgresHost = cast.ToString(coalesce("POSTGRES_HOST", "localhost"))
	config.PostrgresPort = cast.ToString(coalesce("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(coalesce("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("POSTGRES_PASSWORD", "3333"))
	config.PostgresDatabase = cast.ToString(coalesce("POSTGRES_DATABASE", "carpetwash_service"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}
