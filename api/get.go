package main

import (
    "log"
)

type movie struct {
    IMDBId string           `json:"IMDBId"`
    Name string             `json:"Name"`
    Year string             `json:"Year"`
    Score string            `json:"Score"`
    Description string      `json:"Description"`
}

func getMovies() {
    log.Println("GET")
    return
}

func getMovie() (movie) {
    dbMovie := movie{
        "sdf",
        "movie",
        "2000",
        "10",
        "descriptuion",
    }
    return dbMovie
}
