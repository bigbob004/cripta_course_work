package model

type Cache struct {
	UserName               string
	UserID                 int
	Questions              []string
	AuthQuestionToAnswer   map[string]string
	ModelQuestion          []Question
	QuestionsOfChoosenUser []Question
	ChoosenUser            *ChoosenUserInformation

	IsChoosenUserBlocked     bool
	ShuffleQuestions         []string
	RemainingCountAttempts   int
	CountOfRequiredQuestions int
	CountOfInvalidAttempts   int
	CountOfQuestions         int
}

type ChoosenUserInformation struct {
	User      User
	Questions []Question
}
