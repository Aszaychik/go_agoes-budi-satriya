package routers

import (
	"golang_middleware/controllers"
	"golang_middleware/middleware"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo, bc controllers.BookController) {
	booksGroup := e.Group("books")

	booksGroup.POST("", middleware.Validate(bc.CreateBook()))
	booksGroup.GET("", middleware.Validate(bc.GetBooks()))
	booksGroup.GET("/:id", middleware.Validate(bc.GetBook()))
	booksGroup.PUT("/:id", middleware.Validate(bc.UpdateBook()))
	booksGroup.DELETE("/:id", middleware.Validate(bc.DeleteBook()))
}