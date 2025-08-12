package api

import (
	"net/http"
)

type Handlers struct {
	User   *UserHandler
	Health *HealthHandler
	//plan on adding more handlers here depending on the /resources i define
}

func Setup(h *Handlers) http.Handler {
	mux := http.NewServeMux()

	//health endpoint
	mux.HandleFunc("GET /health", h.Health.Check)

	//user endpoints
	mux.HandleFunc("POST /api/v1/users", h.User.CreateUser)

	return mux
}
