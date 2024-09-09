package functions

import (
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	Input    string
	Banner string
	Result string
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

	path := "./common/static/"
	UserData.Input, UserData.Banner = r.Form.Get("input"), r.Form.Get("banner")
	bannerFile := ReadFontFile(path + UserData.Banner + ".txt")
	fontParse := ParseFont(bannerFile, UserData.Banner)
	UserData.Result = "\n" + GenerateAsciiArt(UserData.Input, fontParse)
	length := len(UserData.Result)
	w.Header().Set("Content-Length",strconv.Itoa(length))
	http.Redirect(w, r, "/", http.StatusFound)
}
