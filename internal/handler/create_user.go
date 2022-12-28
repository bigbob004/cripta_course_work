package handler

import (
	"cripta_course_work/internal/model"
	"github.com/sirupsen/logrus"
	"html/template"
	"strconv"

	"net/http"
)

type ViewDataCreateUser struct {
	IsThereUserWithUserNameErr  bool
	CountOfRequiredQuestionsErr bool
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		userName := r.FormValue("user_name")
		//Проверяем, что user_name не занят
		user, err := h.services.Authorization.GetUserByUserName(userName)
		if err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if user != nil {
			tmpl, err := template.ParseFiles("./template/add_user.html")
			if err != nil {
				logrus.Error(err)
			}
			err = tmpl.Execute(w, ViewDataCreateUser{IsThereUserWithUserNameErr: true, CountOfRequiredQuestionsErr: false})
			if err != nil {
				logrus.Error(err)
			}
			return
		}
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
		if countOfRequiredQuestions > len(questions) {
			tmpl, err := template.ParseFiles("./template/add_user.html")
			if err != nil {
				logrus.Error(err)
			}
			err = tmpl.Execute(w, ViewDataCreateUser{IsThereUserWithUserNameErr: false, CountOfRequiredQuestionsErr: true})
			if err != nil {
				logrus.Error(err)
			}
			return
		}
		questionsAndAnswers := []model.Question{}
		for i := 0; i < len(questions); i++ {
			questionsAndAnswers = append(questionsAndAnswers, model.Question{Title: questions[i], Answer: answers[i]})
		}
		newUser := model.User{
			UserName:                 userName,
			CountOfQuestions:         len(questionsAndAnswers),
			CountOfRequiredQuestions: countOfRequiredQuestions,
			CountOfInvalidAttempts:   countOfInvalidAttempts,
		}
		h.services.CreateUser(newUser, questionsAndAnswers)
		http.Redirect(w, r, "/account", http.StatusFound)
	}
}

func (h *Handler) CreateUserView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./template/add_user.html")
		if err != nil {
			logrus.Error(err)
		}
		err = tmpl.Execute(w, ViewDataCreateUser{IsThereUserWithUserNameErr: false, CountOfRequiredQuestionsErr: false})
		if err != nil {
			logrus.Error(err)
		}
	}

}

type signInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
