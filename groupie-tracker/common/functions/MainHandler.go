package functions

import (
	"log"
	"net/http"
	"html/template"
)

type Data struct {
	Artists string
	Locations string
	Dates string
	Relation string
} 

var UserData Data

func MainHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Error(w, "404", http.StatusNotFound)
	}

	if r.Method != "GET" {
		log.Fatal("method not allowed")
	}

	t, err := template.ParseFiles("./common/static/index.html")

	if err != nil {
		log.Fatal("error parsing template")
	}
	t.Execute(w, UserData)
}