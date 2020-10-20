package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func getMovieFromOmdb(id string) (movie, error) {
    apiKey := "e1843a60"
    requestUrl := "http://omdbapi.com/?apikey=" + apiKey + "&plot=full&i="
    //TODO send api request for id for movie and add to db
    resp,err := http.Get(requestUrl + id)
    if err != nil {
        log.Print(err.Error())
        return movie{}, errors.New("The api fucked up")
    }

    b,_ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    var omdbmovie omdbMovie
    err = json.Unmarshal(b, &omdbmovie)

    if omdbmovie.ImdbID == "" {
        return movie{}, errors.New("This movie does not exist")
    }

    putMovieInDatabase(movie)

    return movie{
        omdbmovie.ImdbID,
        omdbmovie.Title,
        omdbmovie.Year,
        omdbmovie.ImdbRating,
        omdbmovie.Plot,
    }, nil
}
