package database

import (
	"github.com/gofrs/uuid"
)

// Dog model properties
type Dog struct {
	DogID    uuid.UUID `gorm:"primary_key"` // This is the Dog Model PK field
	DogName  string
	DogOwner uuid.UUID
}
