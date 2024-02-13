package main

import (
	"net/http"
	"context"
	"os"
	"github.com/unclejoeyb/gorouter/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}




func main() {
	
	component.Render(context.Background(), os.Stdout)
}

