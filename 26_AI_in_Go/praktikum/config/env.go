package config

import (
	"ai_in_go/helper"
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

func NewEnvConfig() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("Error loading .env file: %v", err.Error())
	}
	openAPIKey, err := setOpenAPIKey()
	return openAPIKey, nil
}

func setOpenAPIKey() (string, error){
	openAPIKey := helper.GetEnv("OPEN_API_KEY", "")
	if openAPIKey == "" {
		return openAPIKey, fmt.Errorf("Error loading .env file: %v", errors.New("OPEN_API_KEY is empty"))
	}

	return openAPIKey, nil
}