// Package entities holds all the entities that are shared across all subdomains
package entity

import "github.com/google/uuid"

// Item represents a Item for all sub domains
type Item struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the Item
	Name string
	// Description is the description of the Item
	Description string
}
