package handler

import (
	"cripta_course_work/internal/model"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type AccountViewData struct {
	UserName string
	Users    []model.User
}

func (h *Handler) AccountView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := AccountViewData{UserName: h.cache.UserName}
		if h.cache.UserName == "admin" {
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
