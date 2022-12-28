package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func (h *Handler) Exit(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("Завершение работы приложения")
	os.Exit(0)
}
