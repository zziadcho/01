package main

import (
	"01/ascii-art/common/functions"
	"log"
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", functions.MainHandler)
	http.HandleFunc("/ascii-art-web", functions.HandleAscii)
	fmt.Println("http://localhost" + PORT)
	go functions.AutoLaunchBrowser("http://localhost:8080")
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
