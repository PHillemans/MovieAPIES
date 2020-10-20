package main

import (
	"errors"
	"log"
	"strings"
)

func getMovies() {
    log.Println("GET")
    return
}

func getMovie(requestURI string) (movie, error) {
    var imdbId string
    imdbId = strings.Split(requestURI, "movies/")[1]

    if !isImdbId(imdbId) {
        err := errors.New("Not a imdbId")
        return movie{}, err
    }

    dbMovie := getMovieFromDb(imdbId)

    if dbMovie.IMDBId == "" {
        omdbMovie, err := getMovieFromOmdb(imdbId)
        if err != nil {
            return movie{}, err
        }
        return omdbMovie, nil
    }

    return dbMovie, nil
}

