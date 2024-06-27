package functions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("./common/static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(UserData)
	t.Execute(w, UserData)
}
