package service

import "cripta_course_work/internal/model"

func (s *AuthService) UpdateUser(user model.User) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
