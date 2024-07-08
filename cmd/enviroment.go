package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	// .env  (lee el archivo en compilable por defecto)
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func validateEnviroments() error {
	if os.Getenv("DB_SERVER") == "" {
		return fmt.Errorf("the DB_SERVER env var is mandatory")
	}
	if os.Getenv("ALLOWED_ORIGINS") == "" {
		return fmt.Errorf("the ALLOWED_ORIGINS env var is mandatory")
	}
	if os.Getenv("ALLOWED_METHODS") == "" {
		return fmt.Errorf("the ALLOWED_METHODS env var is mandatory")
	}
	if os.Getenv("IMAGES_DIR") == "" {
		return fmt.Errorf("the IMAGES_DIR env var is mandatory")
	}
	if os.Getenv("JWT_SECRET_KEYS") == "" {
		return fmt.Errorf("the JWT_SECRET_KEYS env var is mandatory")
	}

	//Database
	if os.Getenv("DB_USER") == "" {
		return fmt.Errorf("the DB_USER env var is mandatory")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return fmt.Errorf("the DB_PASSWORD env var is mandatory")
	}
	if os.Getenv("DB_HOST") == "" {
		return fmt.Errorf("the DB_HOST env var is mandatory")
	}
	if os.Getenv("DB_PORT") == "" {
		return fmt.Errorf("the DB_PORT env var is mandatory")
	}
	if os.Getenv("DB_NAME") == "" {
		return fmt.Errorf("the DB_NAME env var is mandatory")
	}
	if os.Getenv("DB_SSL_MODE") == "" {
		return fmt.Errorf("the DB_SSL_MODE env var is mandatory")
	}

	//Paypal envs.
	if os.Getenv("WEBHOOK_ID") == "" {
		return fmt.Errorf("the WEBHOOK_ID env var is mandatory")
	}
	if os.Getenv("VALIDATION_URL") == "" {
		return fmt.Errorf("the VALIDATION_URL env var is mandatory")
	}
	if os.Getenv("CLIENT_ID") == "" {
		return fmt.Errorf("the CLIENT_ID env var is mandatory")
	}
	if os.Getenv("SECRET_ID") == "" {
		return fmt.Errorf("the SECRET_ID env var is mandatory")
	}
	return nil
}
