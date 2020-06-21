type UserUseCase interface {
	getUsers() (user, error)
	createUser() (user, error)
}

type userUseCase struct {

}

func (u *userUseCase) getUsers() (user, error) {

}

func (u *userUseCase) createUser() (user, error) {

}