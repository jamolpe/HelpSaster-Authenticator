package core

import "authorization-service/pkg/models"

// UserRepository : expected methods of the repository
type UserRepository interface {
	SaveUser(user models.User) error
	GetUserByEmail(user models.User) (models.User, error)
}
