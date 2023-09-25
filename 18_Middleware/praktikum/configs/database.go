package configs

import (
	"golang_middleware/helpers"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	DB_Username string
  DB_Password string
  DB_Port     string
  DB_Host     string
  DB_Name     string
}

func InitDBConfig() *DBConfig {
	res := &DBConfig{}
	res = loadConfig()

	return res
}

func loadConfig() *DBConfig {
	res := &DBConfig{}
	err := godotenv.Load()
	if err != nil {
    logrus.Error("Config : Cannot load config file,", err.Error())
		return nil
  }

  res.DB_Username = helpers.GetEnv("DB_USERNAME", "root")
  res.DB_Password = helpers.GetEnv("DB_PASSWORD", "")
  res.DB_Port = helpers.GetEnv("DB_PORT", "3306")
  res.DB_Host = helpers.GetEnv("DB_HOST", "localhost")
  res.DB_Name = helpers.GetEnv("DB_NAME", "orm_code_structure")

	return res
}