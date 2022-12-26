package model

type User struct {
	UserID                   int    `db:"user_id"`
	UserName                 string `db:"user_name"`
	CountOfRequiredQuestions int    `db:"count_of_required_questions"`
	CountOfQuestions         int    `db:"count_of_questions"`
	CountOfInvalidAttempts   int    `db:"count_of_invalid_attempts"`
	IsBlocked                bool   `db:"is_blocked"`
}
