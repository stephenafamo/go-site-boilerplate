package main

import (
    "net/http"
)

func main() {
	http.Handle("/", routes())
    http.ListenAndServe(":80", nil)
}