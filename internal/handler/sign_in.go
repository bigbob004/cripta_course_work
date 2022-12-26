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
		//TODO: валидация
		if h.cache != nil && h.cache.UserName == userName {
			http.Redirect(w, r, fmt.Sprintf("/auth?user_id=%d", h.cache.UserID), http.StatusMovedPermanently)
			return
		}
		user, err := h.services.Authorization.GetUserByUserName(userName)
		if err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		if user != nil {
			//TODO пока не будем париться и будем делать по-тупому
			//TODO КОСТЫЛИ ЕБУЧИЕ, ЗА ЭТО ТЫ ПОПАДЁШЬ В АД!
			if user.IsBlocked {
				data := ViewData{IsThereUserWithUserName: false, IsUserBlocked: true}
				tmpl, _ := template.ParseFiles("./template/login.html")
				tmpl.Execute(w, data)
				return
			}
			h.cache = &model.Cache{UserName: user.UserName, UserID: user.UserID, RemainingCountAttempts: user.CountOfInvalidAttempts, CountOfRequiredQuestions: user.CountOfRequiredQuestions}
			http.Redirect(w, r, fmt.Sprintf("/auth?user_id=%d", user.UserID), http.StatusMovedPermanently)
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
		//TODO: здесь же можно обрабатывать случай IsThereError = true
		data := ViewData{IsThereUserWithUserName: false, IsUserBlocked: false}
		tmpl, _ := template.ParseFiles("./template/login.html")
		tmpl.Execute(w, data)
	}
}
