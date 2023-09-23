package routers

import (
	"intro_echo_golang/controllers"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("", controllers.GetUsersController)
	userGroup.POST("", controllers.CreateUserController)
	userGroup.PUT("/:id", controllers.UpdateUserController)
	userGroup.GET("/:id", controllers.GetUserController)
	userGroup.DELETE("/:id", controllers.DeleteUserController)
}