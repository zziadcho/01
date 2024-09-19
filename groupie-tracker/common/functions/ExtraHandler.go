package functions

import (
	"01/groupie-tracker/common/utils"
	"html/template"
	"net/http"
)

func ExtraHandler(w http.ResponseWriter, r *http.Request) {
	/* protection */
	if r.URL.Path != "/extra" {
		http.Error(w, "403 forbidden", http.StatusForbidden)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	/* variables */
	var relations   utils.Relations
	var locations   utils.Locations
	var dates       utils.Dates

	/* data fetching  */
	err := FetchData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		http.Error(w, "500 error fetching locations data", http.StatusInternalServerError)
		return
	}

	err = FetchData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		http.Error(w, "500 error fetching dates data", http.StatusInternalServerError)
		return
	}

	err = FetchData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		http.Error(w, "500 error fetching relations data", http.StatusInternalServerError)
		return
	}

	/* executing template */
	pageData := utils.PageData{
		Title:     "Artists Info",
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}
	t, err := template.ParseFiles("./common/static/extra.html")
	if err != nil {
		http.Error(w, "500 failed to execute extra template", http.StatusInternalServerError)
	}

	t.Execute(w, pageData)

}
