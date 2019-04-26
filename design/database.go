package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("Owners", func() {
	Description("Storage for hosts and their dogs")
	Store("postgres", gorma.Postgres, func() {
		Description("Postgres database")
		Model("Host", func() {
			Description("Host model properties")
			NoAutomaticIDFields()
			NoAutomaticTimestamps()
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				Description("This is the Host Model PK field")
			})
			Field("age", gorma.Integer, func() {})
			Field("name", gorma.String, func() {})
			RendersTo(HostsResponse)
			HasMany("Dogs", "Dog")
		})
		Model("Dog", func() {
			Description("Dog model properties")
			NoAutomaticIDFields()
			NoAutomaticTimestamps()
			Field("id", gorma.UUID, func() {
				PrimaryKey()
				Description("This is the Dog Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("owner", gorma.UUID, func() {})
			RendersTo(DogMedia)
		})
	})
})
