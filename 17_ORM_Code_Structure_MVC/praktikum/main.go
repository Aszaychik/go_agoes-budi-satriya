package main

import (
	"orm_code_structure/configs"
	"orm_code_structure/controllers"
	"orm_code_structure/models"
	"orm_code_structure/routers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func main() {
	e := echo.New()
	config := configs.InitDBConfig()

	db := models.InitModel(*config)
	models.InitialMigration(db)

	userModel := models.UsersModel{}
	userModel.Init(db)

	userController := controllers.UserController{}
	userController.InitUserController(userModel)

	routers.UserRoutes(e, userController)

	e.Logger.Fatal(e.Start(":8080"))
}