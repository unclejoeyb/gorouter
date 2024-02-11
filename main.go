package main

import (
	"net/http"

	"github.com/unclejoeyb/gorouter/api"
)



func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

