package core

// UserRepository : expected methods of the repository
type UserRepository interface {
	SaveUser(user User) error
	GetUserByEmail(user User) (User, error)
}
