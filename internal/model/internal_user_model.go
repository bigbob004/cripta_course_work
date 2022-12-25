package model

type ShitUser struct {
	UserName                 string
	Questions                []string
	AuthQuestionToAnswer     map[string]string
	ModelQuestion            []Question
	ShuffleQuestions         []string
	RemainingCountAttempts   int
	CountOfRequiredQuestions int
}
