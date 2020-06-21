package handler

import (
	"fmt"
	"server/usecase"
)

type UserHandler interface {
	createUser
	getUsers
}

// UserHandlerを満たす
type userHandler struct {
	userUseCase usecase.UserUseCase
}

// usecase層に対しての処理を行う
func (h *userHandler) func createUser(c echo.Context) error {
	user, err := h.userUseCase.createUser()
	if err := nil {
		fmt.Println(err)
	}
	return c.String("createUser")
}

// usecase層に対しての処理を行う
func (h *userHandler) func getUsers(c echo.Context) error {
	userUseCase.getUsers()
	return c.String("getUsers")
}