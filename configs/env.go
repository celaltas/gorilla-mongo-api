package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func EnvMongoUri() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return "", errors.New("You must set your 'MONGODB_URI' environmental variable.")
	}
	return uri, nil
}
