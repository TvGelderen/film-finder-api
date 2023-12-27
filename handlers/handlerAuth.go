package handlers

import (
	"encoding/json"
	"fmt"
    "net/http"
    "time"
	"github.com/google/uuid"
    "github.com/TvGelderen/film-finder-api/internal/database"
    "github.com/TvGelderen/film-finder-api/internal/auth"
)

func (apiCfg *ApiConfig) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	passwordHash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error hashing password: %v", err))
	}

    _, err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	respondWithJSON(w, 201, "User successfully created")
}

func (apiCfg *ApiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Invalid email or password")
	}

	correctPassword := auth.CheckPasswordWithHash(params.Password, user.PasswordHash)

	if !correctPassword {
		respondWithError(w, 401, "Invalid email or password")
		return
	}

    token, err := auth.CreateNewJWT(user.ID, user.Name)
    if err != nil {
        respondWithError(w, 400, "Failed to create JWT")
    }

    auth.SetToken(w, token)
    
	respondWithJSON(w, 200, mapDbUserToReturnUser(user))
}
