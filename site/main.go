package main

import (
	"net/http"
	"path/to/your/app/models"
)

func main() {
	defer models.DB.Close()
	router := customRouter()
	http.Handle("/", routes(router))
	http.ListenAndServe(":80", nil)
}