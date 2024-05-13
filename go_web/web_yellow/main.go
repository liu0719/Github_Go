package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	video, err := os.Open("./40345471a.webm")
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()

	http.ServeContent(w, r, "40345471a.webm", time.Now(), video)
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":80", nil)
}
