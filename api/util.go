package main

import (
    "net/http"
    "database/sql"
    "log"
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
