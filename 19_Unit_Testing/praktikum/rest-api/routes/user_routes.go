package routes

import (
	"unit_testing/controller"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)
	usersGroup.GET("/:id", userController.GetUserController)
}