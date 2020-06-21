package router

import (
	"github.com/labstack/echo"
	"server/presenter/handler"
)

func newRouter (e *echo.Echo, h handler.IndexHandler) {
	e.GET("/api/v1/users", h.getUsers)
	e.POST("/api/v1/user", h.createUser)
}