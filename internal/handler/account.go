package handler

import (
	"cripta_course_work/internal/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
)

type AccountViewData struct {
	UserName                 string
	CountOfInvalidAttempts   int
	CountOfRequiredQuestions int
	Users                    []model.User
	Err                      string
	IsThereError             bool
}

func (h *Handler) AccountView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var isThereError bool
		errUpdating := r.FormValue("err")
		if len(errUpdating) != 0 {
			isThereError = true
		}
		data := AccountViewData{UserName: h.cache.UserName, CountOfRequiredQuestions: h.cache.CountOfRequiredQuestions, CountOfInvalidAttempts: h.cache.CountOfInvalidAttempts, IsThereError: isThereError, Err: errUpdating}
		if h.cache.UserName == "admin" {
			if h.cache.ChoosenUser != nil {
				h.cache.ChoosenUser = nil
			}
			users, err := h.services.GetAllUsers()
			if err != nil {
				logrus.Error("handler/account/AccountView: ", err)
				w.Write([]byte("что-то пошло не так"))
				return
			}
			data.Users = users
		}

		tmpl, err := template.ParseFiles("./template/account.html")
		if err != nil {
			logrus.Error(err)
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (h *Handler) EditAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		countOfRequiredQuestions, _ := strconv.Atoi(r.FormValue("count_of_required_questions"))
		countOfInvalidAttempts, _ := strconv.Atoi(r.FormValue("count_of_invalid_attempts"))

		if countOfRequiredQuestions > len(h.cache.ModelQuestion) {
			err := "Количество обязательных вопросов не можеть быть большего общего количества вопросов"
			http.Redirect(w, r, fmt.Sprintf("/account?err=%s", err), http.StatusFound)
			return
		}
		h.cache.CountOfRequiredQuestions = countOfRequiredQuestions
		h.cache.CountOfInvalidAttempts = countOfInvalidAttempts

		user := model.User{UserID: h.cache.UserID, UserName: h.cache.UserName, CountOfInvalidAttempts: countOfInvalidAttempts, CountOfRequiredQuestions: countOfRequiredQuestions}

		err := h.services.UpdateUser(user)
		if err != nil {
			logrus.Error("handler/EditAccount: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		http.Redirect(w, r, "/account", http.StatusFound)
	}
}
