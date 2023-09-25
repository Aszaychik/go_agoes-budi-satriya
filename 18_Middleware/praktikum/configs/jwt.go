package configs

import (
	"golang_middleware/helpers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AuthConfig struct {
	Secret        []byte
}

type JWTClaim struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func InitAuthConfig() *AuthConfig {
	res := &AuthConfig{}
	res = loadAuthConfig()

	return res
}

func loadAuthConfig() *AuthConfig {
	res := &AuthConfig{}
	err := godotenv.Load()
	if err != nil {
    logrus.Error("Config : Cannot load config file,", err.Error())
		return nil
  }
	
	res.Secret = []byte(helpers.GetEnv("SECRET", "secret"))

	return res
}