package main

import (
	"github.com/TvGelderen/film-finder-api/internal/database"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, mapDbUserToUser(user))
}
