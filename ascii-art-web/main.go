package main

import (
	"log"
	"net/http"
	"01/ascii-art/common/functions"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", functions.MainHandler)
	http.HandleFunc("/ascii-art-web", functions.HandleAscii)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
