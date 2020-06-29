package handler

import (
	"fmt"
	"server/usecase"
)

type UserHandler interface {
	createUser(c echo.Context) error
	getUsers(c echo.Context) error
}

// UserHandlerを満たす
type userHandler struct {
	userUseCase usecase.UserUseCase
}

// usecase層に対しての処理を行う
func (h *userHandler) func createUser(c echo.Context) error {
	user, err := h.userUseCase.createUser(c)
	if err := nil {
		fmt.Println(err)
	}
	return c.JSON(user)
}

// usecase層に対しての処理を行う
func (h *userHandler) func getUsers(c echo.Context) error {
	user, err := h.userUseCase.getUsers(c)
	if err := nil {
		fmt.Println(err)
	}
	return c.JSON(user)
}