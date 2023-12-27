package handlers

import (
    "fmt"
    "net/http"
    "github.com/TvGelderen/film-finder-api/internal/database"
    "github.com/TvGelderen/film-finder-api/internal/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString, err := auth.GetToken(r)
        if err != nil {
            respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
            return
        }

        id, err := auth.GetIdFromJWT(tokenString)
        if err != nil {
            respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
            return
        }

        user, err := apiCfg.DB.GetUserById(r.Context(), id)
        if err != nil {
            respondWithError(w, 400, fmt.Sprintf("Unable to get user: %v", err))
            return
        }

        handler(w, r, user)
    }
}
