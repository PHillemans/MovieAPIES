package main

import (
    "net/http"
    "log"
    "encoding/json"
)

func main() {
    var port string = ":8080"
    log.Printf("\n\nStarting server on port: %v\n", port)

    // single movie
    http.HandleFunc("/movies", moviesHandler)
    http.HandleFunc("/movies/", movieHandler)

    http.ListenAndServe(port, nil)
}

func moviesHandler(w http.ResponseWriter, req *http.Request) {
    // @TODO fix cors and make sending get requests work
    enableCors(&w)
    switch req.Method {
        case "GET":
            getMovies()
            w.Write([]byte("something"))

        case "POST":
            postMovies()

        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            w.Write([]byte("Try doing a POST or a GET request"))
    }
}

func movieHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "GET":
        movie := getMovie()
        json.NewEncoder(w).Encode(movie)

    case "POST":
        postMovie()

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Try doing a POST or a GET request"))
    }
}

