package service

import (
	"cripta_course_work/internal/model"
	"cripta_course_work/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User, questions []model.Question) error
	GetUserByUserName(userName string) (*model.User, error)
	GetQuestionsByUserID(userID int) ([]model.Question, error)
	ValidateQuestions(userAnswers []string, user *model.Cache) bool
	GetAllUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	UpdateQuestion(question model.Question) error
	DropQuestion(questionID int) error
	AddQuestion(question model.Question) (int, error)
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
