package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func Config(key, path string) (string, error) {
	err := godotenv.Load(path)
	if err != nil {
		return "", errors.New("wrong env path")
	}

	envKey := os.Getenv(key)
	if envKey == "" {
		return "", errors.New("wrong env variable name")
	}

	return envKey, nil
}
