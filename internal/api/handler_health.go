package api

import (
	"net/http"

	"github.com/05blue04/Flow/internal/config"
)

type HealthHandler struct {
	app *config.App
}

func NewHealthHandler(app *config.App) *HealthHandler {
	return &HealthHandler{
		app: app,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
