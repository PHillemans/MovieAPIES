package main

import (
	"log"
	"strings"
)

func getMovies() {
    log.Println("GET")
    return
}

func getMovie(requestURI string) (movie) {
    var imdbId string
    imdbId = strings.Split(requestURI, "movies/")[1]

    db := getDatabase();

    query := "SELECT * FROM movies WHERE imdbid is ?;"

    prep, _ := db.Prepare(query)
    result := prep.QueryRow(imdbId)

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

    if dbMovie.IMDBId == "" {
        dbMovie = getMovieFromOmdb(requestURI);
    }

    return dbMovie
}

