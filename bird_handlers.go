package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {

	// Get data from db store
	birds, err := store.GetBirds()

	// Convert bird list to json
	birdListBytes, err := json.Marshal(birds)

	// Print error to console
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write json list to response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {

	// Create new instance
	bird := Bird{}

	// Parse html
	err := r.ParseForm()

	// Display error
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get bird info from form
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Add bird to database
	err = store.CreateBird(&bird)

	if err != nil {
		fmt.Println(err)
	}

	// Redirect to original html
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
