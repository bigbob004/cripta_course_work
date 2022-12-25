package db

import (
	"cripta_course_work/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	usersTable     = "users"
	questionsTable = "questions"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_name, count_of_required_questions, count_of_questions, count_of_invalid_attempts) values(?, ?, ?, ?) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.CountOfRequiredQuestions, user.CountOfQuestions, user.CountOfInvalidAttempts)
	logrus.Debugf("repository/db/GetUser: %v", query)
	if err := row.Scan(&id); err != nil {
		logrus.Error("repository/db/CreateUser: ", err)
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username string) ([]model.User, error) {
	var users []model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_name=$1", usersTable)
	err := r.db.Select(&users, query, username)
	logrus.Debugf("repository/db/GetUser: %v", query)
	if err != nil {
		logrus.Error("repository/db/GetUser: ", err)
		return []model.User{}, err
	}

	return users, err
}

func (r *AuthPostgres) CreateQuestions(questions []model.Question) (int, error) {
	var id int

	valueStrings := make([]string, 0, len(questions))
	valueArgs := make([]interface{}, 0, len(questions)*3)
	for _, question := range questions {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, question.UserID)
		valueArgs = append(valueArgs, question.Title)
		valueArgs = append(valueArgs, question.Answer)
	}
	query := fmt.Sprintf("INSERT INTO questions (user_id, question_text, answer_text) VALUES %s",
		strings.Join(valueStrings, ","))
	logrus.Debugf("repository/db/GetUser: %v", query)
	_, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		logrus.Error("repository/db/CreateQuestion: ", err)
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetAllUsers() ([]model.User, error) {
	var users []model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_name !='admin'", usersTable)
	logrus.Debugf("repository/db/GetAllUsers: %v", query)
	err := r.db.Select(&users, query)
	if err != nil {
		logrus.Error("repository/db/GetAllUsers: ", err)
		return []model.User{}, err
	}

	return users, err
}

func (r *AuthPostgres) GetAllQuestions() ([]model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT * FROM %s", questionsTable)
	logrus.Debugf("repository/db/GetAllQuestionsr: %v", query)
	err := r.db.Select(&questions, query)
	if err != nil {
		logrus.Error("repository/db/GetAllQuestions: ", err)
		return []model.Question{}, err
	}

	return questions, err
}

func (r *AuthPostgres) GetQuestionsByUserID(userID int) ([]model.Question, error) {
	var questions []model.Question
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", questionsTable)
	logrus.Debugf("repository/db/GetQuestionsByUser: %v", query)
	err := r.db.Select(&questions, query, userID)
	if err != nil {
		logrus.Error("repository/db/GetQuestionsByUser: ", err)
		return []model.Question{}, err
	}

	return questions, err
}
