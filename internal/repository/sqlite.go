package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
)

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func NewSqlite() (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB

	if !checkFileExists("auth.db") {
		db, err = sqlx.Open("sqlite3", "auth.db")
		if err != nil {
			logrus.Fatalf("faied to initialize db: %s", err.Error())
			return nil, err
		}
		_, err = db.Exec(`create table users
	   (
	       user_id            integer primary key autoincrement not null,
	       user_name          text    not null,
	       count_of_required_questions integer not null,
	       count_of_questions integer not null,
	       count_of_invalid_attempts integer not null
	   	   constraint user_name_unique unique (user_name)
	   );
	   create table questions
	   (
	       question_id   INTEGER not null
	           primary key  autoincrement,
	       user_id       INTEGER not null
	           references users,
	       question_text TEXT    not null,
	       answer_text   TEXT    not null
	   );
	   insert into users
	   (user_name, count_of_required_questions, count_of_questions, count_of_invalid_attempts)
	   values ("admin", 1, 1, 3);
	   insert into questions
	   (user_id, question_text, answer_text)
	   values
	   (1, "Вопрос первого входа админа", "123");`)

		if err != nil {
			return nil, err
		}
	} else {
		db, err = sqlx.Open("sqlite3", "auth.db")
		if err != nil {
			logrus.Fatalf("faied to initialize db: %s", err.Error())
			return nil, err
		}
	}
	return db, nil
}
