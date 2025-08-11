package router

import (
	"net/http"

	"github.com/05blue04/Flow/internal/handlers"
)

type Handlers struct {
	User   *handlers.UserHandler
	Health *handlers.HealthHandler
}

func Setup(h *Handlers) http.Handler {
	mux := http.NewServeMux()

	//health endpoint
	mux.HandleFunc("GET /health", h.Health.Check)

	//user endpoints
	mux.HandleFunc("POST /api/v1/users", h.User.CreateUser)

	return mux
}
