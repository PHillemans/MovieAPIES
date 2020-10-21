package main

import (
    "net/http"
    "log"
)

func main() {
    createTable()

    var port string = ":8080"
    log.Printf("\n\nStarting server on port: %v\n", port)

    // single movie
    http.HandleFunc("/movies", moviesHandler)
    http.HandleFunc("/movies/", movieHandler)

    http.ListenAndServe(port, nil)
}

func moviesHandler(w http.ResponseWriter, req *http.Request) {
    setHeaders(&w) //TODO: Remove this on prod
    switch req.Method {
        case "GET":
            log.Println(req.URL.Query())
            movies := getMovies()
            writeMoviesResponse(w, movies)

        case "POST":
            postMovies()

        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            w.Write([]byte("Try doing a POST or a GET request"))
    }
}

func movieHandler(w http.ResponseWriter, req *http.Request) {
    setHeaders(&w) //TODO: Remove this on prod
    switch req.Method {
    case "GET":
        movie, err := getMovie(req.RequestURI)
        if err != nil {
            w.WriteHeader(http.StatusUnprocessableEntity)
            writeErrorResponse(w, err)
            return
        }
        writeMovieResponse(w, movie)

    case "POST":
        postMovie()

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("Try doing a POST or a GET request"))
    }
}

func insertMovieInToDatabase(newMovie movie) {
    db := getDatabase()
    defer db.Close()

    query := "INSERT INTO MOVIES (imdbid, name, year, score, desc, poster) VALUES (?, ?, ?, ?, ?, ?)"
    prep, err := db.Prepare(query)
    if err != nil {
        log.Println(err.Error())
        return
    }

    prep.Exec(newMovie.IMDBId, newMovie.Name, newMovie.Year, newMovie.Score, newMovie.Description, newMovie.Poster)
}

func createTable() {
    db := getDatabase()
    defer db.Close()
    query := `
        Create TABLE IF NOT EXISTS movies
            (
                id INTEGER PRIMARY KEY,
                imdbid TEXT,
                name TEXT,
                year TEXT,
                score TEXT,
                desc TEXT,
                poster TEXT
            )`

    prep, err := db.Prepare(query)
    if err != nil {
        log.Fatalln("db error:", err.Error())
    }
    prep.Exec()
}
