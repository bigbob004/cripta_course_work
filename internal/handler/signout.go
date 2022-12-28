package handler

import "net/http"

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.cache = nil
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
