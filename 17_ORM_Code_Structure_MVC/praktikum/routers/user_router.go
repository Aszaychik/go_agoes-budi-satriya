package routers

import (
	"orm_code_structure/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controllers.UserController) {
	usersGroup := e.Group("users")

	usersGroup.GET("", uc.GetUsers())
	usersGroup.GET("/:id", uc.GetUser())
	usersGroup.POST("", uc.CreateUser())
	usersGroup.PUT("/:id", uc.UpdateUser())
	usersGroup.DELETE("/:id", uc.DeleteUser())
}