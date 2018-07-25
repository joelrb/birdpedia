package main

import (
	"database/sql"
)

// Interface to database
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

// Implement interface, takes db connection
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES ($1,$2)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {

	// Query database
	rows, err := store.db.Query("SELECT species, description from birds")

	// Return error if any
	if err != nil {
		return nil, err
	}

	// Defer closing
	defer rows.Close()

	// Create data structure
	birds := []*Bird{}

	// Itirate record
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		bird := &Bird{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		birds = append(birds, bird)
	}

	return birds, nil

}

var store Store

// Initialize
func InitStore(s Store) {
	store = s
}
