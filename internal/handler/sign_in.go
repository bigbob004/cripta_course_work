package handler

import (
	"cripta_course_work/internal/model"
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	IsThereUserWithUserName bool
	IsUserBlocked           bool
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		userName := r.FormValue("user_name")

		if h.cache != nil && h.cache.UserName == userName {
			http.Redirect(w, r, fmt.Sprintf("/auth?user_id=%d", h.cache.UserID), http.StatusFound)
			return
		}
		user, err := h.services.Authorization.GetUserByUserName(userName)
		if err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		if user != nil {
			if user.IsBlocked {
				data := ViewData{IsThereUserWithUserName: false, IsUserBlocked: true}
				tmpl, _ := template.ParseFiles("./template/login.html")
				tmpl.Execute(w, data)
				return
			}
			h.cache = &model.Cache{UserName: user.UserName, UserID: user.UserID, RemainingCountAttempts: user.CountOfInvalidAttempts, CountOfInvalidAttempts: user.CountOfInvalidAttempts, CountOfRequiredQuestions: user.CountOfRequiredQuestions, CountOfQuestions: user.CountOfQuestions}
			http.Redirect(w, r, fmt.Sprintf("/auth?user_id=%d", user.UserID), http.StatusFound)
		} else {
			data := ViewData{IsThereUserWithUserName: true, IsUserBlocked: false}
			tmpl, _ := template.ParseFiles("./template/login.html")
			tmpl.Execute(w, data)
		}
	}
}

func (h *Handler) SignInView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := ViewData{IsThereUserWithUserName: false, IsUserBlocked: false}
		tmpl, _ := template.ParseFiles("./template/login.html")
		tmpl.Execute(w, data)
	}
}
