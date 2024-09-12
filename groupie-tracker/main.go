package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// const PORT = ":8080"

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, _ := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	responseData, _ := io.ReadAll(response.Body)

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(responseObject.Pokemon)
}

// fmt.Println("http://localhost" + PORT)
// http.HandleFunc("/", functions.MainHandler)
// err := http.ListenAndServe(PORT, nil)
// if err != nil {
// 	log.Fatal(err)
// 	return
// }
