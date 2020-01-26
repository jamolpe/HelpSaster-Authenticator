package core


func (s *userSrv) findIfUserExist(user *User) bool {
	expectedUser, _ := s.repo.GetUserByEmail(*user)
	if expectedUser != (User{}) {
		return true
	}
	return false
}
