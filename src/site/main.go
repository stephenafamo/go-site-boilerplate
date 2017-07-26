package main

import (
	"net/http"
)

func main() {
	router := customRouter()
	http.Handle("/", routes(router))
	http.ListenAndServe(":80", nil)
}
