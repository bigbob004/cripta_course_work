package handler

import (
	"cripta_course_work/internal/model"
	tools "cripta_course_work/pkg"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
)

type ChoosenUserView struct {
	User         model.User
	Err          string
	IsThereError bool
}

func (h *Handler) EditUserView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		var isThereError bool
		errUpdating := r.FormValue("err")
		if len(errUpdating) != 0 {
			isThereError = true
		}

		userID, _ := strconv.Atoi(r.FormValue("user_id"))
		//Тут записывает в кэш пользователя, которого собираемся редактировать
		users, err := h.services.GetUserByID(userID)
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if len(users) == 0 {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		user := users[0]
		questions, err := h.services.GetQuestionsByUserID(user.UserID)
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if len(questions) == 0 {
			logrus.Error(fmt.Sprintf("handler/EditUserView: для юзера с именем %s нет вопросов в БД", user.UserName))
			w.Write([]byte("что-то пошло не так"))
			return
		}
		h.cache.ChoosenUser = &model.ChoosenUserInformation{User: user, Questions: questions}
		tmpl, _ := template.ParseFiles("./template/user.html")
		tmpl.Execute(w, ChoosenUserView{User: h.cache.ChoosenUser.User, IsThereError: isThereError, Err: errUpdating})
	}
}

func (h *Handler) EditUserInformation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		is_blocked := tools.ConvertToBool(r.FormValue("is_blocked"))
		countOfRequiredQuestions, _ := strconv.Atoi(r.FormValue("count_of_required_questions"))
		countOfInvalidAttempts, _ := strconv.Atoi(r.FormValue("count_of_invalid_attempts"))

		if countOfRequiredQuestions > len(h.cache.ChoosenUser.Questions) {
			err := "Количество обязательных вопросов не можеть быть большего общего количества вопросов"
			http.Redirect(w, r, fmt.Sprintf("/edit_user?err=%s&user_id=%d", err, h.cache.ChoosenUser.User.UserID), http.StatusFound)
			return
		}

		//TODO Валидация на количество обязательных вопросов
		h.cache.ChoosenUser.User.IsBlocked = is_blocked
		h.cache.ChoosenUser.User.CountOfRequiredQuestions = countOfRequiredQuestions
		h.cache.ChoosenUser.User.CountOfInvalidAttempts = countOfInvalidAttempts

		err := h.services.UpdateUser(h.cache.ChoosenUser.User)
		if err != nil {
			logrus.Error("handler/EditUserInformation: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		http.Redirect(w, r, "/account", http.StatusFound)
	}
}
