package functions

import (
	"01/groupie-tracker/common/utils"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	/* protection */
	if r.URL.Path != "/" {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var artists []utils.Artists
	
	/* data fetching */
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		http.Error(w, "500 failed to fetch artists data", http.StatusInternalServerError)
		return
	}

	/* template execute */
	t, err := template.ParseFiles("./common/static/index.html")
	if err != nil {
		http.Error(w, "500 failed to execute index template", http.StatusInternalServerError)
	}

	t.Execute(w, artists) 

}
