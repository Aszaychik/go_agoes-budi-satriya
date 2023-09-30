package routes

import (
	"unit_testing/controller"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.CreateUserController)
	usersGroup.POST("/login", userController.LoginUserController)
}