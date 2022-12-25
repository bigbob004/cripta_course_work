package handler

import (
	"html/template"
	"net/http"
)

func (h *Handler) AccountView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, _ := template.ParseFiles("./template/account.html")
		tmpl.Execute(w, h.questionOfCurrentUser)
	}
}
