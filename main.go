package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the static folder containing index.html
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
