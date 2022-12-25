package handler

import "net/http"

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.questionOfCurrentUser = nil
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}
