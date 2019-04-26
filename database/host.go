package database

import (
	"github.com/gofrs/uuid"
)

// Host model properties
type Host struct {
	ID   uuid.UUID `gorm:"primary_key"` // This is the Host Model PK field
	Age  int
	Name string
}

type HostWithDogs struct {
	ID      uuid.UUID `gorm:"primary_key"` // This is the Host Model PK field
	Name    string
	Age     int
	DogID   uuid.UUID
	DogName string
}
