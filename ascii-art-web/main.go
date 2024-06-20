package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Data struct {
	Art string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("getting info")
		fmt.Println("get")
	} else if r.Method == http.MethodPost {
		log.Println("posting info")
		fmt.Println("post")
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := Data{
		Art: "Ziad",
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", FormHandler)
	fmt.Println("server started on port :8080")
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
