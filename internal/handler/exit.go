package handler

import (
	"log"
	"net/http"
)

func (h *Handler) Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Завершение работы приложения")
}
