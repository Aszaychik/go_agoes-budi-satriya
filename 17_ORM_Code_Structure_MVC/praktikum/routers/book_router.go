package routers

import (
	"orm_code_structure/controllers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo, bc controllers.BookController) {
	booksGroup := e.Group("books")

	booksGroup.GET("", bc.GetBooks())
	booksGroup.GET("/:id", bc.GetBook())
	booksGroup.POST("", bc.CreateBook())
	booksGroup.PUT("/:id", bc.UpdateBook())
	booksGroup.DELETE("/:id", bc.DeleteBook())
}