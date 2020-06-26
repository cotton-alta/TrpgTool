package handler

import (
	"testing"
	"net/http/httptest"
	"github.com/labstack/echo"
)

type mocUserUseCase struct{}

func (u *mocUserUseCase) getUsers() ([]*model.User, error) {

}

func (u *mocUserUseCase) createUser() ([]*model.User, error) {

}

func TestUserHandler_getUsers(t *testing.T) {
	usecase := &mocUserUseCase{}
	h := newHandler(usercase)

	e := echo.New()
	req := httptest.NewRequest(echo.Get, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
}