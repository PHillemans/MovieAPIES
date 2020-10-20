package main

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

func setHeaders(w *http.ResponseWriter) {
    enableCors(w)
    (*w).Header().Set("Content-Type","appication/json")
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getDatabase() *sql.DB {
    db, err := sql.Open("sqlite3", "watchlist.db")

    if err != nil {
        log.Fatalln("Couldn't open database", err)
    }
    return db
}

func isImdbId(id string) (bool) {
    match,_ := regexp.Match("tt[0-9]+", []byte(id))
    return match
}

func getMovieFromDb(requestId string) (movie) {
    db := getDatabase();

    query := "SELECT * FROM movies WHERE imdbid is ?;"

    prep, _ := db.Prepare(query)
    result := prep.QueryRow(requestId)

    var (
        id string
        imdbid string
        movieName string
        year string
        score string
        description string
    )

    result.Scan(&id, &imdbid, &movieName, &year, &score, &description)

    dbMovie := movie{
        imdbid, movieName, year, score, description,
    }
    return dbMovie
}

func putMovieInDatabase(movie omdbMovie) {

    //TODO put into database
}
