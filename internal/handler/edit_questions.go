package handler

import (
	"cripta_course_work/internal/model"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

type ViewDataEditQuestions struct {
	ModelQuestions []model.Question
}

func (h *Handler) EditQuestionsView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		funcMap := template.FuncMap{
			// The name "inc" is what the function will be called in the template text.
			"inc": func(i int) int {
				return i + 1
			},
		}
		name := path.Base("./template/question_redactor_new.html")
		tmpl, err := template.New(name).Funcs(funcMap).ParseFiles("./template/question_redactor_new.html")
		if err != nil {
			logrus.Error(err)
		}
		err = tmpl.Execute(w, h.questionOfCurrentUser.ModelQuestion)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (h *Handler) EditQuestions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		//TODO сделать отдельную функцию сбора данных
		//TODO: маппинг и валидация
		userName := r.FormValue("user_name")
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
		user := model.User{Username: userName, CountOfInvalidAttempts: countOfInvalidAttempts}
		questionsAndAnswers := make([]model.Question, 0, len(questions))
		for i := 0; i < len(questions); i++ {
			questionsAndAnswers = append(questionsAndAnswers, model.Question{Title: questions[i], Answer: answers[i]})
		}

		h.services.CreateUser(user, questionsAndAnswers)
		http.Redirect(w, r, "/add_user", http.StatusMovedPermanently)
	}
}
