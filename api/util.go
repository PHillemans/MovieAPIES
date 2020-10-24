package main

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
    defer db.Close()

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
        poster string
    )

    result.Scan(&id, &imdbid, &movieName, &year, &score, &description, &poster)

    dbMovie := movie{
        imdbid, movieName, year, score, description, poster,
    }
    return dbMovie
}

func movieCount() (int) {
    db := getDatabase()
    defer db.Close()

    query := "SELECT COUNT(imdbid) FROM movies"
    prep,_ := db.Prepare(query)
    result := prep.QueryRow()
    var count int
    result.Scan(&count)
    return count
}

func getMoviesFromDbWithOffset(offset int) ([]movie) {
    db := getDatabase()
    defer db.Close()

    query := "SELECT * FROM movies ORDER BY id DESC LIMIT "+ strconv.Itoa(offset) +",10;"

    prep,err := db.Prepare(query)
    if err != nil {
        log.Println(err.Error())
        log.Fatalln("prep")
    }

    var dbMovies []movie

    results, err := prep.Query()
    if err != nil {
        log.Fatalln(err.Error())
    }
    for results.Next() {
        var (
            id string
            imdbid string
            name string
            year string
            score string
            description string
            poster string
        )
        results.Scan (&id, &imdbid, &name, &year, &score, &description, &poster)
        dbMovie := movie{
            imdbid,
            name,
            year,
            score,
            description,
            poster,
        }
        dbMovies = append(dbMovies, dbMovie)
    }

    return dbMovies
}

func getMoviesFromDb() ([]movie) {
    db := getDatabase()
    defer db.Close()

    query := "SELECT * FROM movies"

    prep,_ := db.Prepare(query)
    var dbMovies []movie
    results,_ := prep.Query()
    for results.Next() {
        var (
            id string
            imdbid string
            name string
            year string
            score string
            description string
            poster string
        )
        results.Scan (&id, &imdbid, &name, &year, &score, &description, &poster)
        dbMovie := movie{
            imdbid,
            name,
            year,
            score,
            description,
            poster,
        }
        dbMovies = append(dbMovies, dbMovie)
    }

    return dbMovies
}
