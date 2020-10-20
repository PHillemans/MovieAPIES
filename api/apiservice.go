package main

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, movie movie) {
    if movie.IMDBId == "" {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Could not find the movie you were looking for."))
        return
    }

    json.NewEncoder(w).Encode(movie)
}
