package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user recognition.User) (int, error)
	GetUser(username, password string) (recognition.User, error)
}

type Repoistory struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repoistory {
	return &Repoistory{
		Authorization: NewAuthPostgres(db),
	}
}
