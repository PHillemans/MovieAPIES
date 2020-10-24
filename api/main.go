package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
    createTable()

    var port string = ":8080"
    log.Printf("\n\nStarting server on port: %v\n", port)

    // single movie
    http.HandleFunc("/movies", moviesHandler)
    http.HandleFunc("/movies/", movieHandler)
    http.HandleFunc("/movieamount", amountHandler)

    http.ListenAndServe(port, nil)
}

func amountHandler(w http.ResponseWriter, req *http.Request) {
    setHeaders(&w)
    if (req.Method != "GET") {
        return
    }
    amount := strconv.Itoa(getMovieAmount())
    w.Write([]byte(amount))
}

func moviesHandler(w http.ResponseWriter, req *http.Request) {
    setHeaders(&w) //TODO: Remove this on prod
    switch req.Method {
        case "GET":
            page := req.URL.Query()["page"][0]
            var movies []movie
            pageParam,_ := strconv.Atoi(page)
            if len(page) > 1 || pageParam > 0 {
                movies = getMovies(pageParam - 1)
            } else {
                movies = getAllMovies()
            }
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
