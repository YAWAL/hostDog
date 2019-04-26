package database

import (
	"github.com/YAWAL/hostDog/config"
	"github.com/go-pg/pg"
	"github.com/gofrs/uuid"
)

const (
	HostIDField     = "hostID"
	DogIDField      = "dogID"
	GetAllHost      = "SELECT * FROM hosts"
	GetDogsByHostID = "SELECT * FROM dogs WHERE dog_owner = ?"
	GetHostByID     = "SELECT  h.id, h.name, h.age, d.dog_id,d.dog_name FROM hosts h LEFT JOIN dogs d ON h.id = d.dog_owner WHERE h.id = ?"
)

func Postgres(config config.DBConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     config.User,
		Password: config.Pass,
		Database: config.DBName,
	})
	return db
}

func GenerateID() (uuid.UUID, error) {
	return uuid.NewV4()
}

func UUID(input string) (uuid.UUID, error) {
	return uuid.FromString(input)
}
