package handlers

import "net/http"

func HandlerSuccess(w http.ResponseWriter, r *http.Request) {
    respondWithJSON(w, 200, struct{}{})
}
