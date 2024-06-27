package functions

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Art    string
	Banner string
}

var UserData Data

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-web" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	UserData.Art = r.Form.Get("art")
	UserData.Banner = r.Form.Get("banner")
	bannerFile := ReadFontFile(AddTxtExtension(UserData.Banner))
	fmt.Println(bannerFile)
	fmt.Println(UserData)

	http.Redirect(w, r, "/", http.StatusFound)
}
