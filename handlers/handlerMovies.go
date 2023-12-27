package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TvGelderen/film-finder-api/internal/database"
)

func (apiCfg *ApiConfig) HandlerSaveMovie(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		MovieId int32 `json:"movieId"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
	}

    movie, err := apiCfg.DB.SaveMovie(r.Context(), database.SaveMovieParams {
        MovieID: params.MovieId,
        UserID: user.ID,
    })

    respondWithJSON(w, 201, movie)
}

func (apiCfg *ApiConfig) HandlerGetSavedMovies(w http.ResponseWriter, r *http.Request, user database.User) {
    movies, err := apiCfg.DB.GetUserSavedMovies(r.Context(), user.ID)
    if err != nil {
        respondWithError(w, 400, "Error retrieving saved movies")
        return
    }

	respondWithJSON(w, 200, movies)
}

func (apiCfg *ApiConfig) HandlerRemoveMovie(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		MovieId int32 `json:"movieId"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
        return
	}

    err = apiCfg.DB.RemoveMovie(r.Context(), database.RemoveMovieParams{
        MovieID: params.MovieId, 
        UserID: user.ID,
    })
    if err != nil {
        respondWithError(w, 400, "Error removing movie")
        return
    }

	respondWithJSON(w, 200, "Movie removed successfully")
}
