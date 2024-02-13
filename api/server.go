package api

import (
	"encoding/json"
	"net/http"
	"context"
	"os"
	"github.com/unclejoeyb/gorouter/tree/main/api/templates"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

type Item struct {
	ID  uuid.UUID `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Server struct {
	*mux.Router

	shoppingList []Item
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		shoppingList: []Item{},
	}
	s.Routes()
	return s
}

// func (s *Server) Routes() {
// 	s.HandleFunc("/items", s.createItem()).Methods("POST")
// 	s.HandleFunc("/items", s.listItems()).Methods("GET")
// 	s.HandleFunc("/items/{id}", s.deleteItem()).Methods("DELETE")
// 	s.HandleFunc("/items/{id}", s.updateItem()).Methods("PUT")
// }


func (s *Server) createItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item.ID = uuid.New()
		s.shoppingList = append(s.shoppingList, item)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}



func (s *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		

		
	}
}

func (s *Server) deleteItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		found := false
		for i, item := range s.shoppingList {
			if item.ID == id {
				s.shoppingList = append(s.shoppingList[:i], s.shoppingList[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
	}
}

func (s *Server) updateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr, _ := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		found := false
		for i, oldItem := range s.shoppingList {
			if oldItem.ID == id {
				s.shoppingList[i] = item
				found = true
				break
			}
		}
		if !found {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}
	}
}
