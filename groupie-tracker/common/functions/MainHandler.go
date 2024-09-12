package functions

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

type Data struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

var UserData Data

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creation_date"`
	FirstAlbum   string   `json:"first_album"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concert_dates"`
	Relations    string   `json:"relations"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure correct URL path
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Only allow GET method
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fetch artists data
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var responseObject []Artists
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the HTML template
	tmpl := template.Must(template.ParseFiles(filepath.Join("./common/static/index.html")))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "error executing template", http.StatusInternalServerError)
	}
}

