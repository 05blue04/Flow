package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/05blue04/Flow/internal/auth"
	"github.com/05blue04/Flow/internal/config"
	"github.com/05blue04/Flow/internal/database"
	"github.com/google/uuid"
)

type UserHandler struct {
	app *config.App
}

func NewUserHandler(app *config.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respondWithError(w, 400, "couldn't decode parameters", err)
		return
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 500, "couldn't create hash for password", err)
		return
	}

	u, err := h.app.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		Username:       params.Username,
		Email:          params.Email,
		HashedPassword: hash,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})

	if err != nil {
		respondWithError(w, 400, "error creating user", err)
		return
	}

	newU := types.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	respondWithJSON(w, 201, newU)
}
