package main

import (
    "net/http"
)

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Controll-Allow-Origin", "*")
}
