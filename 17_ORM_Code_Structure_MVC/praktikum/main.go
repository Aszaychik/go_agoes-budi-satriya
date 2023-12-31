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

	bookModel := models.BooksModel{}
	bookModel.Init(db)

	blogModel := models.BlogsModel{}
	blogModel.Init(db)

	userController := controllers.UserController{}
	userController.InitUserController(userModel)

	bookController := controllers.BookController{}
	bookController.InitBookController(bookModel)

	blogController := controllers.BlogController{}
	blogController.InitBlogController(blogModel)

	routers.UserRoutes(e, userController)
	routers.BookRoutes(e, bookController)
	routers.BlogRoutes(e, blogController)

	e.Logger.Fatal(e.Start(":8080"))
}