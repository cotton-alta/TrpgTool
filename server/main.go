package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"server/presenter/router"
	"server/presenter/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	i := newHandler()
	router.newRouter(e, i)

	e.Start(":3001")
}
