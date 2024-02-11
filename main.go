package main

import (
	"net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ninja":
		switch r.Method {
		case "GET":
			NinjaHandlers.get(w, r)
		case "POST":
			NinjaHandlers.post(w, r)
		case "PUT":
			NinjaHandlers.put(w, r)
		case "DELETE":
			NinjaHandlers.delete(w, r)
		}
	case "/dojo":
		switch r.Method {
		case "GET":
			DojoHandlers.get(w, r)
		case "POST":
			DojoHandlers.post(w, r)
		case "PUT":
			DojoHandlers.put(w, r)
		case "DELETE":
			DojoHandlers.delete(w, r)
		}
	default:
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}		

func main() {
	http.HandleFunc("/", baseHandler)
	http.ListenAndServe(":8080", nil)

}

