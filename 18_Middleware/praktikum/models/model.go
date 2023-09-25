package models

import (
	"fmt"
	"golang_middleware/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(config configs.DBConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    config.DB_Username,
    config.DB_Password,
    config.DB_Host,
    config.DB_Port,
    config.DB_Name,
  )

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		logrus.Error("Error connecting to database: ", err)
		return nil
	}
	return db
}

func InitialMigration(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &Book{}, &Blog{})
	db.AutoMigrate(&User{}, &Book{})
}
