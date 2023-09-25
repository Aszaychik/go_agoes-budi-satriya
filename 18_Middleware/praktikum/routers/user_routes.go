package routers

import (
	"golang_middleware/controllers"
	"golang_middleware/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, uc controllers.UserController) {
	usersGroup := e.Group("users")

	usersGroup.POST("", uc.CreateUser())
	usersGroup.POST("/login", uc.Login())
	usersGroup.GET("/logout", uc.Logout())
	
	usersGroup.GET("", middleware.Validate(uc.GetUsers()))
	usersGroup.GET("/:id", middleware.Validate(uc.GetUser()))
	usersGroup.PUT("/:id", middleware.Validate(uc.UpdateUser()))
	usersGroup.DELETE("/:id", middleware.Validate(uc.DeleteUser()))
}