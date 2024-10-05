package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/source"
)

const Host, Port = "localhost", ":8080"

// Main function to start the server
func main() {

	// Set up HTTP handler
	http.HandleFunc("/", source.HomePageHandler)
	http.HandleFunc("/artist", source.ArtistDetailsHandler)

	// Set path for static material
	fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	fmt.Println("Server starting on port " + Port + "...")
	fmt.Println("ctrl + click to open: http://" + Host + Port + "/")
	http.ListenAndServe(Host+Port, nil)
}
