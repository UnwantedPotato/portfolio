package main

import (
	"fmt"
	"net/http"
)

func clicked(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p style='color: green;'>Content swapped dynamically by htmx!</p>")
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /", fileServer)
	mux.HandleFunc("GET /clicked", clicked)
	http.ListenAndServe(":8080", mux)
}
