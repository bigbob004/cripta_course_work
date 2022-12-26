package repository

import (
	"cripta_course_work/internal/model"
	db2 "cripta_course_work/internal/repository/db"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string) ([]model.User, error)
	CreateQuestions(questions []model.Question) (int, error)
	GetAllUsers() ([]model.User, error)
	GetAllQuestions() ([]model.Question, error)
	GetQuestionsByUserID(userID int) ([]model.Question, error)
	UpdateUser(user model.User) error
	DropUser(userName string) error
	UpdateQuestion(question model.Question) error
	DropQuestion(questionID int) error
	AddQuestion(question model.Question) (int, error)
}

type Repoistory struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repoistory {
	return &Repoistory{
		Authorization: db2.NewAuthPostgres(db),
	}
}
