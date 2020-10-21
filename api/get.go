package main

import (
	"errors"
	"log"
	"strings"
)

func getMovies() ([]movie) {
    log.Println("GET")
    movies := getMoviesFromDb()
    return movies
}

func getMovie(requestURI string) (movie, error) {
    var imdbId string
    imdbId = strings.Split(requestURI, "movies/")[1]

    if !isImdbId(imdbId) {
        err := errors.New("Not a imdbId")
        return movie{}, err
    }

    // Getting db from the database
    dbMovie := getMovieFromDb(imdbId)

    if dbMovie.IMDBId == "" {
        omdbMovie, err := getMovieFromOmdb(imdbId)
        if err != nil {
            return movie{}, err
        }
        return omdbMovie, nil
    }
    log.Println("got from db")

    return dbMovie, nil
}

