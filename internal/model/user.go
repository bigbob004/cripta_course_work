package model

type User struct {
	UserID                   int    `db:"user_id"`
	Username                 string `db:"user_name"`
	CountOfRequiredQuestions int    `db:"count_of_required_questions"`
	CountOfQuestions         int    `db:"count_of_questions"`
	CountOfInvalidAttempts   int    `db:"count_of_invalid_attempts"`
}
