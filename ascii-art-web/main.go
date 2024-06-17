package main

import (
	"log"
	"net/http"
	"text/template"
)

type Data struct{
	Title string
	Head string
	Content string
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "Welcome",
		Head: "This is the head",
		Content: "This is my content",
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
}

func cssHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "style.css")
}
func main() {
	http.HandleFunc("/", handler)
    http.HandleFunc("/styles.css", cssHandler)
	log.Println("server is on port :8080")
	http.ListenAndServe(":8080", nil)
}
