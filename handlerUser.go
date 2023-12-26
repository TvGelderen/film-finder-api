package main

import (
	"net/http"
    "github.com/TvGelderen/film-finder-api/internal/database"
)

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, mapDbUserToUser(user))
}
