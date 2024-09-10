package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

func LoadConfig() error {
    err := godotenv.Load("./configs/.env")

    if err != nil {
     return fmt.Errorf("Error loading .env: %w", err)
    }

    return err
}

func Get(key string) string {
    return os.Getenv(key)
}