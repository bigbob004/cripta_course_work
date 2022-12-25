package model

type Question struct {
	QuestionID int    `db:"question_id" json:"question_id"`
	UserID     int    `db:"user_id"  json:"user_id"`
	Title      string `db:"question_text"  json:"title"`
	Answer     string `db:"answer_text"  json:"answer"`
}
