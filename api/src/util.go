package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
    "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func importData() {
    db := getDatabase()
    defer db.Close()

    csvFile, err := os.Open("../watchlist.csv")
    if err != nil {
        log.Println(err)
        return
    }

    r := csv.NewReader(csvFile)
    if _, err := r.Read(); err != nil {
        log.Println("ERRRROR terror")
    }

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }

        imdbId := record[1]
        name := record[5]
        year,_ := strconv.ParseFloat(record[10], 64)
        score,_ := strconv.Atoi(record[8])
        description := ""

        newMovie := movie{
            imdbId,
            name,
            score,
            year,
            description,
        }
        insertMovieInToDatabase(newMovie)
    }

}

func getDatabase() *sql.DB {
    db, err := sql.Open("sqlite3", "watchlist.db")

    if err != nil {
        log.Fatalln("Couldn't open database", err)
    }
    return db
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
                year int,
                score float,
                desc TEXT
            )`

    prep, err := db.Prepare(query)
    if err != nil {
        log.Fatalln("db error:", err.Error())
    }
    prep.Exec()
}
