package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// HomeHandler -
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World.")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT should be set.")
	}

	http.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":"+port, nil)
}
