package service

import (
	"cripta_course_work/internal/model"
	tools "cripta_course_work/pkg"
	"sort"
)

// Костыли и сюда добрались

func (s *AuthService) ValidateQuestions(userAnswers []string, user *model.ShitUser) bool {
	//Сначала отбираем ответы на обязательные вопросы
	authAnswers := make([]string, 0, len(user.AuthQuestionToAnswer))
	authUserAnswers := make([]string, 0, len(user.AuthQuestionToAnswer))
	for index, item := range user.ShuffleQuestions {
		if value, ok := user.AuthQuestionToAnswer[item]; ok {
			authUserAnswers = append(authUserAnswers, userAnswers[index])
			authAnswers = append(authAnswers, value)
		}
	}
	//Сравниваем ответы пользователя и реальные ответы
	sort.Strings(authUserAnswers)
	sort.Strings(authAnswers)
	return tools.IsEq(authUserAnswers, authAnswers)
}
