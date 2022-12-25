package service

import (
	"cripta_course_work/internal/model"
	"cripta_course_work/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User, questions []model.Question) error
	SignIn(userName string) (*model.User, error)
	GetQuestionsByUserID(userID int) ([]model.Question, error)
	ValidateQuestions(userAnswers []string, user *model.ShitUser) bool
	//GenerateToken(username string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repoistory) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
