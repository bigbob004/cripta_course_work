package handler

import (
	"cripta_course_work/internal/model"
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	IsThereError           bool
	RemainingCountAttempts int
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		userName := r.FormValue("user_name")
		//TODO: валидация
		user, err := h.services.Authorization.SignIn(userName)
		if err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		if user != nil {
			//TODO пока не будем париться и будем делать по-тупому
			//TODO КОСТЫЛИ ЕБУЧИЕ, ЗА ЭТО ТЫ ПОПАДЁШЬ В АД!
			h.questionOfCurrentUser = &model.ShitUser{UserName: user.Username, RemainingCountAttempts: user.CountOfInvalidAttempts, CountOfRequiredQuestions: user.CountOfRequiredQuestions}
			http.Redirect(w, r, fmt.Sprintf("/auth?user_id=%d", user.UserID), http.StatusMovedPermanently)
		} else {
			data := ViewData{IsThereError: true}
			tmpl, _ := template.ParseFiles("./template/login.html")
			tmpl.Execute(w, data)
		}
	}
}

func (h *Handler) SignInView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//TODO: здесь же можно обрабатывать случай IsThereError = true
		data := ViewData{IsThereError: false}
		tmpl, _ := template.ParseFiles("./template/login.html")
		tmpl.Execute(w, data)
	}
}
