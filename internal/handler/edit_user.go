package handler

import (
	tools "cripta_course_work/pkg"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

/*
func (h *Handler) EditUserView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		userName := r.FormValue("user_name")
		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		//Тут записывает в кэш пользователя, которого собираемся редактировать
		user, err :=h.services.GetUserByUserName(userName)
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		questions, err := h.services.GetQuestionsByUserID(userID)
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if len(questions) == 0 {
			logrus.Error("handler/EditUserView: ", fmt.Errorf("у выбранного пользователя не созданы контрольные вопросы"))
			w.Write([]byte("что-то пошло не так"))
			return
		}
		h.cache.QuestionsOfChoosenUser = questions

		if
		tmpl, _ := template.ParseFiles("./template/user.html")
		tmpl.Execute(w, userName)
	}
}

*/

func (h *Handler) EditUserView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		userName := r.FormValue("user_name")
		//Тут записывает в кэш пользователя, которого собираемся редактировать
		user, err := h.services.GetUserByUserName(userName)
		if err != nil {
			logrus.Error("handler/EditUserView: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if user == nil {
			logrus.Error(fmt.Sprintf("handler/EditUserView: юзера с именем %s нет в БД", userName))
			w.Write([]byte("что-то пошло не так"))
			return
		}
		h.cache.ChoosenUser = *user

		tmpl, _ := template.ParseFiles("./template/user.html")
		tmpl.Execute(w, userName)
	}
}

func (h *Handler) BlockUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		is_blocked := tools.ConvertToBool(r.FormValue("is_blocked"))
		if h.cache.ChoosenUser.IsBlocked != is_blocked {
			h.cache.ChoosenUser.IsBlocked = is_blocked
			err := h.services.UpdateUser(h.cache.ChoosenUser)
			if err != nil {
				logrus.Error("handler/BlockUser: ", err)
				w.Write([]byte("что-то пошло не так"))
				return
			}
		}
		http.Redirect(w, r, "/account", http.StatusMovedPermanently)
	}
}
