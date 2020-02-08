package core

import "authorization-service/pkg/models"

func (s *userSrv) findIfUserExist(user *models.User) bool {
	expectedUser, _ := s.repo.GetUserByEmail(*user)
	if expectedUser != (models.User{}) {
		return true
	}
	return false
}

func (s *userSrv) getUserFromDatabase(user models.User) *models.User {
	var expectedUser models.User
	expectedUser, _ = s.repo.GetUserByEmail(user)
	return &expectedUser
}
