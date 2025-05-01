package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"real-time-forum/_database"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	if err := database.InitDB(); err != nil {
		fmt.Printf("error creating the database: %v", err)
		return
	}

	http.HandleFunc("/api/posts", getPostsHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server has started http://localhost:%v/", port)
	fmt.Println(http.ListenAndServe(":" + port, nil))
}

func getPostFromDB() ([]Post, error) {
	rows, err := database.DB.Query("SELECT id, title, content FROM PostTable LIMIT 50")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := getPostFromDB()

	if err != nil {
		fmt.Printf("error getting posts from the database: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
