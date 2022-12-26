package service

import "cripta_course_work/internal/model"

func (s *AuthService) UpdateQuestion(question model.Question) error {
	err := s.repo.UpdateQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) DropQuestion(questionID int) error {
	err := s.repo.DropQuestion(questionID)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) AddQuestion(question model.Question) (int, error) {
	questionID, err := s.repo.AddQuestion(question)
	if err != nil {
		return 0, err
	}
	return questionID, nil
}
