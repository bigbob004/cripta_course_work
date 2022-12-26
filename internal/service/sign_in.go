package service

import "cripta_course_work/internal/model"

func (s *AuthService) GetUserByUserName(userName string) (*model.User, error) {
	users, err := s.repo.GetUser(userName)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}
