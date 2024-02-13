package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestServer(t *testing.T) {
	s := NewServer()
	if s == nil {
		t.Error("NewServer returned nil")
	}
}

func TestServerRoutes(t *testing.T) {
	s := NewServer()
	if s == nil {
		t.Error("NewServer returned nil")
	}
	s.Routes()
}
func TestCreateItem(t *testing.T) {
	s := NewServer()

	// Create a request with a sample item
	item := Item{Name: "Test Item", Description: "This is a test item description"}
	body, _ := json.Marshal(item)
	req := httptest.NewRequest("http.MethodPost", "/items", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the createItem handler function
	handler := http.HandlerFunc(s.createItem())
	handler.ServeHTTP(rr, req)
	

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	var responseItem Item
	err := json.Unmarshal(rr.Body.Bytes(), &responseItem)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Verify the item ID is not empty
	if responseItem.ID.String() == "" {
		t.Error("Item ID is empty")
	}

	// Verify the item is added to the server's shopping list
	if len(s.shoppingList) != 1 {
		t.Errorf("Expected shopping list length 1, but got %d", len(s.shoppingList))
	}
}
func TestListItems(t *testing.T) {
	s := NewServer()

	// Add some items to the shopping list
	item1 := Item{Name: "Item 1", Description: "Description 1"}
	item2 := Item{Name: "Item 2", Description: "Description 2"}
	s.shoppingList = append(s.shoppingList, item1, item2)

	// Create a request to list items
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	rr := httptest.NewRecorder()

	// Call the listItems handler function
	handler := http.HandlerFunc(s.listItems())
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	var responseItems []Item
	err := json.Unmarshal(rr.Body.Bytes(), &responseItems)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	// Verify the number of items in the response
	expectedItemCount := len(s.shoppingList)
	if len(responseItems) != expectedItemCount {
		t.Errorf("Expected %d items in the response, but got %d", expectedItemCount, len(responseItems))
	}

	// Verify the content of each item in the response
	for i, item := range responseItems {
		expectedItem := s.shoppingList[i]
		if item.Name != expectedItem.Name || item.Description != expectedItem.Description {
			t.Errorf("Expected item %d to have name '%s' and description '%s', but got name '%s' and description '%s'",
				i+1, expectedItem.Name, expectedItem.Description, item.Name, item.Description)
		}
	}
}
func TestDeleteItem(t *testing.T) {
	s := NewServer()

	// Add a sample item to the shopping list
	item := Item{ID: uuid.New(), Name: "Test Item", Description: "This is a test item description"}
	s.shoppingList = append(s.shoppingList, item)
	t.Log(item.ID.String())

	req := httptest.NewRequest("DELETE", "/items/"+item.ID.String(), nil)
	t.Log(req.URL.String())
	rr := httptest.NewRecorder()

	// Call the deleteItem handler function
	handler := http.HandlerFunc(s.deleteItem())
	handler.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}


	// Verify that the item is removed from the shopping list
	found := false
	for _, listItem := range s.shoppingList {
		if listItem.ID == item.ID {
			found = true
			break
		}
	}
	if found {
		t.Error("Item was not removed from the shopping list")
	}


	// Verify that the item is removed from the shopping list
	if len(s.shoppingList) != 0 {
		t.Errorf("Expected shopping list length 0, but got %d", len(s.shoppingList))
	}
}
func TestUpdateItem(t *testing.T) {
	s := NewServer()

	// Add an item to the shopping list
	item := Item{ID: uuid.New(), Name: "Test Item", Description: "This is a test item description"}
	s.shoppingList = append(s.shoppingList, item)

	// Create a request with an updated item
	updatedItem := Item{ID: item.ID, Name: "Updated Item", Description: "This is an updated item description"}
	body, _ := json.Marshal(updatedItem)
	req := httptest.NewRequest(http.MethodPut, "/items/"+item.ID.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the updateItem handler function
	handler := http.HandlerFunc(s.updateItem())
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Verify the item is updated in the server's shopping list
	found := false
	for _, listItem := range s.shoppingList {
		if listItem.ID == item.ID {
			if listItem.Name != updatedItem.Name || listItem.Description != updatedItem.Description {
				t.Errorf("Item not updated correctly")
			}
			found = true
			break
		}
	}
	if !found {
		t.Error("Item not found in the shopping list")
	}
}