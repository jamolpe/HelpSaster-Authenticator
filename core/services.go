package core

type UserServiceInterface interface {
	UserRegister(user *User) (bool, error)
	Authenticate(loginUserData LoginUserData)
}

func UserRegister(user *User) (bool, error) {

	return true, nil
}
