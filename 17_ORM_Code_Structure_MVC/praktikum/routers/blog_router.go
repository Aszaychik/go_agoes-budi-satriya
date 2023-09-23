package routers

import (
	"orm_code_structure/controllers"

	"github.com/labstack/echo/v4"
)

func BlogRoutes(e *echo.Echo, bc controllers.BlogController) {
	blogsGroup := e.Group("blogs")

	blogsGroup.GET("", bc.GetBlogs())
	blogsGroup.GET("/:id", bc.GetBlog())
	blogsGroup.POST("", bc.CreateBlog())
	blogsGroup.PUT("/:id", bc.UpdateBlog())
	blogsGroup.DELETE("/:id", bc.DeleteBlog())
}