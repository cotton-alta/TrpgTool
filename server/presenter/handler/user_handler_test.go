package handler

import (
	"testing"
	"net/http/httptest"
	"github.com/labstack/echo"
	"server/domain/model"
)

// handlerのテストであるため下階層のusecaseをモックとして用意する
type mocUserUseCase struct{}

func (u *mocUserUseCase) getUsers() ([]*model.User, error) {
	return getMocUsers(3)
}

func (u *mocUserUseCase) createUser() ([]*model.User, error) {
	return getMocUser(1)
}

func getMocUsers(i int) []*model.User {
	users := []*model.User{}
	for n := 0; n < i; n++ {
		user := getMocUser(n)
		users := append(users, user)
	}
	return users
}

func getMocUser(i int) *model.User {
	user := *model.User{
		id: i,
		name: string(i)
	}
	return user
}

func TestUserHandler_getUsers(t *testing.T) {
	usecase := &mocUserUseCase{}
	h := newHandler(usercase)

	e := echo.New()
	req := httptest.NewRequest(echo.Get, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
}