package handler

import (
	"net/http"
)

func (h *Handler) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.cache == nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
