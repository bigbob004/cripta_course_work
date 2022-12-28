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

type SQL struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *SQL {
	return &SQL{db: db}
}

func (r *SQL) CreateUser(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_name, count_of_required_questions, count_of_questions, count_of_invalid_attempts) values(?, ?, ?, ?) RETURNING user_id", usersTable)
	row := r.db.QueryRow(query, user.UserName, user.CountOfRequiredQuestions, user.CountOfQuestions, user.CountOfInvalidAttempts)
	logrus.Debugf("repository/db/GetUser: %v", query)
	if err := row.Scan(&id); err != nil {
		logrus.Error("repository/db/CreateUser: ", err)
		return 0, err
	}

	return id, nil
}

func (r *SQL) GetUserByID(userID int) ([]model.User, error) {
	var users []model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersTable)
	err := r.db.Select(&users, query, userID)
	logrus.Debugf("repository/db/GetUserByID: %v", query)
	if err != nil {
		logrus.Error("repository/db/GetUserByID: ", err)
		return []model.User{}, err
	}
	return users, err
}

func (r *SQL) GetUser(username string) ([]model.User, error) {
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

func (r *SQL) CreateQuestions(questions []model.Question) (int, error) {
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

func (r *SQL) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_name !='admin'", usersTable)
	logrus.Debugf("repository/db/GetAllUsers: %v", query)
	err := r.db.Select(&users, query)
	if err != nil {
		logrus.Error("repository/db/GetAllUsers: ", err)
		return []model.User{}, err
	}

	return users, err
}

func (r *SQL) GetAllQuestions() ([]model.Question, error) {
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

func (r *SQL) GetQuestionsByUserID(userID int) ([]model.Question, error) {
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

func (r *SQL) UpdateUser(user model.User) error {
	query := fmt.Sprintf("UPDATE %s SET user_name=?, count_of_required_questions=?, count_of_invalid_attempts=?, count_of_questions=?, is_blocked=? WHERE user_id=?", usersTable)
	logrus.Debugf("repository/db/UpdateUser: %v", query)
	_, err := r.db.Exec(query, user.UserName, user.CountOfRequiredQuestions, user.CountOfInvalidAttempts, user.CountOfQuestions, user.IsBlocked, user.UserID)
	if err != nil {
		logrus.Error("repository/db/UpdateUser: ", err)
		return err
	}
	return nil
}

func (r *SQL) UpdateQuestion(question model.Question) error {
	query := fmt.Sprintf("UPDATE %s SET question_text=?, answer_text=? WHERE question_id=?", questionsTable)
	logrus.Debugf("repository/db/UpdateQuestion: %v", query)
	_, err := r.db.Exec(query, question.Title, question.Answer, question.QuestionID)
	if err != nil {
		logrus.Error("repository/db/UpdateQuestion: ", err)
		return err
	}
	return nil
}

func (r *SQL) DropQuestion(questionID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE question_id=?", questionsTable)
	logrus.Debugf("repository/db/DropQuestion: %v", query)
	_, err := r.db.Exec(query, questionID)
	if err != nil {
		logrus.Error("repository/db/DropQuestion: ", err)
		return err
	}
	return nil
}

func (r *SQL) DropUser(userName string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_name=?", usersTable)
	logrus.Debugf("repository/db/DropUser: %v", query)
	_, err := r.db.Exec(query, userName)
	if err != nil {
		logrus.Error("repository/db/DropUser: ", err)
		return err
	}
	return nil
}

func (r *SQL) AddQuestion(question model.Question) (int, error) {
	var questionID int
	query := fmt.Sprintf("INSERT INTO %s (user_id, question_text, answer_text) values (?, ?, ?) returning question_id", questionsTable)
	logrus.Debugf("repository/db/AddQuestion: %v", query)
	err := r.db.QueryRow(query, question.UserID, question.Title, question.Answer).Scan(&questionID)
	if err != nil {
		logrus.Error("repository/db/AddQuestion: ", err)
		return 0, err
	}
	return questionID, nil
}
