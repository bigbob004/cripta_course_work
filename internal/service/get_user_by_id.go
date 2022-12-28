package service

import "cripta_course_work/internal/model"

func (s *AuthService) GetUserByID(userID int) ([]model.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
