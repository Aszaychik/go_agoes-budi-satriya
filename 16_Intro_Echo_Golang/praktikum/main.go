package main

import (
	"intro_echo_golang/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()


	routers.NewUserRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}