package functions

import (
	"log"
	"net/http"
	"strconv"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	formErr := r.ParseForm()
	if formErr != nil {
		log.Fatal(formErr)
		return
	}

	var fileName string
	if r.Form.Get("fileName") == "" {
		fileName = "Download"
	} else {
		fileName = r.Form.Get("fileName")
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+".txt")
	w.Header().Set("Content-Type", "text/plain")
	fileContent := UserData.Result
	length := len(fileContent)
	w.Header().Set("Content-Length", strconv.Itoa(length))

	_, err := w.Write([]byte(fileContent))
	if err != nil {
		http.Error(w, "Error generating file", http.StatusInternalServerError)
		return
	}

}
