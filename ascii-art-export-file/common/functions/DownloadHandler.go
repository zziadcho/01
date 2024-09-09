package functions

import (
	"fmt"
	"net/http"
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

	w.Header().Set("Content-Disposition", "attachment; filename=download.txt")

	fileContent := UserData.Result
	size := r.Header.Get("Content-Length")
	fmt.Println(size)
	_, err := w.Write([]byte(fileContent))
	if err != nil {
		http.Error(w, "Error generating file", http.StatusInternalServerError)
		return
	}

}
