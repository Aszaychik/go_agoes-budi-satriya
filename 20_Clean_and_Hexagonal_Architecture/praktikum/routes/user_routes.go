package routes

import (
	"simple_clean_code/controller"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, userController controller.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", userController.CeateUserController())
	usersGroup.GET("", userController.GetAllUserController())
}