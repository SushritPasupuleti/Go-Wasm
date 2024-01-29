package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    // Serve static files from the 'dist' directory
    dir := http.Dir("../dist")
    r.Handle("/*", http.FileServer(dir))

    // Start the server
    http.ListenAndServe(":5000", r)
}
