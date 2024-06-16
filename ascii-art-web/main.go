package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"log"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	var generatedArt string
	if r.Method == http.MethodGet {
		log.Println("get")
	} else if r.Method == http.MethodPost {
		log.Println("post")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error", http.StatusBadRequest)
			return
		}
		input := r.Form.Get("input")
		if input != "" {
			log.Println("post request with input:", input)
		}
		printableSplit := functions.ArgSplitter(input)
		generatedArt = functions.GeneratorLoop(printableSplit, functions.ParseFont(functions.ReadFontFile("standard.txt"), "standard"))
		// temp, err := template.New("index.html").Parse(generatedArt)

		w.Header().Set("Content-Type", "text/html")
	}
	fmt.Fprint(w, "<div class='container'><pre>"+generatedArt+"</pre></div>")
	http.ServeFile(w, r, "index.html")
}
func main() {
	http.HandleFunc("/", FormHandler)
	log.Println("server is on port :8080")
	http.ListenAndServe(":8080", nil)
}
