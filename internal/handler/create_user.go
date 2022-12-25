package handler

import (
	"cripta_course_work/internal/model"
	"strconv"

	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		//TODO сделать отдельную функцию сбора данных
		//TODO: маппинг и валидация
		userName := r.FormValue("user_name")
		countOfRequiredQuestions, _ := strconv.Atoi(r.FormValue("count_of_required_questions"))
		countOfInvalidAttempts, _ := strconv.Atoi(r.FormValue("count_of_invalid_attempts"))
		var questions []string
		var answers []string
		for key, values := range r.PostForm {
			if key == "question" {
				questions = values
			} else if key == "answer" {
				answers = values
			}
		}
		questionsAndAnswers := []model.Question{}
		for i := 0; i < len(questions); i++ {
			questionsAndAnswers = append(questionsAndAnswers, model.Question{Title: questions[i], Answer: answers[i]})
		}
		user := model.User{
			Username:                 userName,
			CountOfQuestions:         len(questionsAndAnswers),
			CountOfRequiredQuestions: countOfRequiredQuestions,
			CountOfInvalidAttempts:   countOfInvalidAttempts,
		}
		h.services.CreateUser(user, questionsAndAnswers)
		http.Redirect(w, r, "/account", http.StatusMovedPermanently)
	case http.MethodGet:
		http.Redirect(w, r, "/send_new_user", http.StatusMovedPermanently)
	}
}

type signInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
