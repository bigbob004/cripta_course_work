package model

type Cache struct {
	UserName                 string
	UserID                   int
	Questions                []string
	AuthQuestionToAnswer     map[string]string
	ModelQuestion            []Question
	QuestionsOfChoosenUser   []Question
	ChoosenUser              User
	IsChoosenUserBlocked     bool
	ShuffleQuestions         []string
	RemainingCountAttempts   int
	CountOfRequiredQuestions int
}
