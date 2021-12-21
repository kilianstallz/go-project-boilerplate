package domain

import "github.com/google/uuid"

// The Domain Object implementation

type Idea struct {
	ID          uuid.UUID
	Name        string
	Description string
}
