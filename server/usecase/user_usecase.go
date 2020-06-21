type UserUseCase interface {
	getUsers
	createUser
}

type userUseCase struct {

}

func (u *userUseCase) getUsers() (user, error) {

}

func (u *userUseCase) createUser() (user, error) {
	
}