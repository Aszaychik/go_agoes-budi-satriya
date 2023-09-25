package main

import (
	"golang_middleware/configs"
	"golang_middleware/controllers"
	"golang_middleware/models"
	"golang_middleware/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// blogModel := models.BlogsModel{}
	// blogModel.Init(db)

	userController := controllers.UserController{}
	userController.InitUserController(userModel)

	bookController := controllers.BookController{}
	bookController.InitBookController(bookModel)

	// blogController := controllers.BlogController{}
	// blogController.InitBlogController(blogModel)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routers.UserRoutes(e, userController)
	routers.BookRoutes(e, bookController)
	// routers.BlogRoutes(e, blogController)

	e.Logger.Fatal(e.Start(":8080"))
}