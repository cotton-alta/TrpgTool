package usecase

import (
	"context"
)

type UserUseCase interface {
	getUsers(ctx context.Context) ([]*model.User, error)
	createUser(ctx context.Context) ([]*model.User, error)
}

type userUseCase struct {

}

func (u *userUseCase) getUsers() ([]*model.User, error) {

}

func (u *userUseCase) createUser() ([]*model.User, error) {

}