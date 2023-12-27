package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TvGelderen/film-finder-api/internal/auth"
	"github.com/TvGelderen/film-finder-api/internal/database"
	"github.com/google/uuid"
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
        if strings.Contains(err.Error(), "users_email_key") {
            respondWithError(w, 400, fmt.Sprintf("That email is already taken"))
            return
        }
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
        fmt.Printf("error: %v\n", err)
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 401, "Invalid email or password")
        return
	}

	correctPassword := auth.CheckPasswordWithHash(params.Password, user.PasswordHash)

	if !correctPassword {
		respondWithError(w, 401, "Invalid email or password")
		return
	}

    token, err := auth.CreateNewJWT(user.ID, user.Name)
    if err != nil {
        respondWithError(w, 400, "Failed to create JWT")
        return
    }

    auth.SetToken(w, token)
    
	respondWithJSON(w, 200, mapDbUserToReturnUser(user))
}

func (apiCfg *ApiConfig) HandlerLogout(w http.ResponseWriter, r *http.Request, user database.User) {
    auth.RemoveToken(w)
}
