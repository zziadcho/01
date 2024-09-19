package functions

import (
	"01/groupie-tracker/common/utils"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "error getting info from API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "error reading data", http.StatusInternalServerError)
		return
	}

	var responseObject []utils.Artists
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		http.Error(w, "error parsing data", http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./common/static/index.html")
	if err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}

	t.Execute(w, responseObject)

}
