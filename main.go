package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"handlers/handlers"
)

func main() {
	// Use the handler function for routing
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/search", handlers.IndexHandler)

	log.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
