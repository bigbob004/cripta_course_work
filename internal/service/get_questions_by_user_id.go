package service

import "cripta_course_work/internal/model"

func (s *AuthService) GetQuestionsByUserID(userID int) ([]model.Question, error) {
	questions, err := s.repo.GetQuestionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
