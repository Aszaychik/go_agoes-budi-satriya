package main

import (
	"simple_clean_code/config"
	"simple_clean_code/controller"
	"simple_clean_code/model/domain"
	"simple_clean_code/repository"
	"simple_clean_code/routes"
	"simple_clean_code/service"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.User{})

	userRepository := repository.NewUserRepository(db)
	userServcice := service.NewUserService(userRepository)
	userController := controller.NewUserController(userServcice)

	routes.NewUserRoutes(app, userController)


	app.Logger.Fatal(app.Start(":8080"))
}
