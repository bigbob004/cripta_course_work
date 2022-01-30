package service

import (
	"github.com/bigbob004/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user recognition.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repoistory) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
