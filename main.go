package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The Music Forecast")
		// http.ServeFile(w, r, 'client/index.html')
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}