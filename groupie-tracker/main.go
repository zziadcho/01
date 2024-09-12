package main

import (
	"01/groupie-tracker/common/functions"
	"log"
	"net/http"
)

const PORT = ":8080"

// type Response struct {
// 	Name    string    `json:"name"`
// 	Pokemon []Pokemon `json:"pokemon_entries"`
// }

// type Pokemon struct {
// 	EntryNo int            `json:"entry_number"`
// 	Species PokemonSpecies `json:"pokemon_species"`
// }

// type PokemonSpecies struct {
// 	Name string `json:"name"`
// }
func main() {
	// Serve static files (CSS, JS, images, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register the main handler
	http.HandleFunc("/", functions.MainHandler)

	// Start the server
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

	