package authorizationcore

import "authorization-service/pkg/models"

// UserRepository : expected methods for user repository
type UserRepository interface {
	SaveUser(user models.User) error
	GetUserByEmail(user models.User) (models.User, error)
}
