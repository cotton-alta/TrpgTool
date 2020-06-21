package handler

import (
	"github.com/labstack/echo"
	"server/usecase"
)

type IndexHandler interface {
	UserHandler
}

// IndexHandlerを満たす
type indexHandler struct {
	userHandler
}

// indexHandlerのコンストラクタ
func newHandler() indexHandler {
	return &indexHandler{}
}