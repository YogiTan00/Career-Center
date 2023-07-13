package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DEBUG                  string
	TIMEOUT                string
	ADDRESS                string
	DB_HOST                string
	DB_PORT                string
	DB_USER                string
	DB_PASS                string
	DB_NAME                string
	PATH_IMAGE_UPLOAD      string
	PATH_IMAGE_UPLOAD_META string
	PATH_FILE_UPLOAD       string
	PATH_FILE_UPLOAD_META  string
	CONFIG_SMTP_HOST       string
	CONFIG_SMTP_PORT       string
	CONFIG_SENDER_NAME     string
	CONFIG_AUTH_EMAIL      string
	CONFIG_AUTH_PASSWORD   string
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func ConfigEnv() Config {
	return Config{
		ADDRESS:                goDotEnvVariable("ADDRESS"),
		DB_HOST:                goDotEnvVariable("DB_HOST"),
		DB_PORT:                goDotEnvVariable("DB_PORT"),
		DB_USER:                goDotEnvVariable("DB_USER"),
		DB_PASS:                goDotEnvVariable("DB_PASS"),
		DB_NAME:                goDotEnvVariable("DB_NAME"),
		CONFIG_SMTP_HOST:       goDotEnvVariable("CONFIG_SMTP_HOST"),
		CONFIG_SMTP_PORT:       goDotEnvVariable("CONFIG_SMTP_PORT"),
		CONFIG_SENDER_NAME:     goDotEnvVariable("CONFIG_SENDER_NAME"),
		CONFIG_AUTH_EMAIL:      goDotEnvVariable("CONFIG_AUTH_EMAIL"),
		CONFIG_AUTH_PASSWORD:   goDotEnvVariable("CONFIG_AUTH_PASSWORD"),
		PATH_IMAGE_UPLOAD:      goDotEnvVariable("PATH_IMAGE_UPLOAD"),
		PATH_IMAGE_UPLOAD_META: goDotEnvVariable("PATH_IMAGE_UPLOAD_META"),
		PATH_FILE_UPLOAD:       goDotEnvVariable("PATH_FILE_UPLOAD"),
		PATH_FILE_UPLOAD_META:  goDotEnvVariable("PATH_FILE_UPLOAD_META"),
	}
}
