package main

import (
	"github.com/YAWAL/hostDog/app"
	"github.com/YAWAL/hostDog/database"
	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
)

// DogController implements the dog resource.
type DogController struct {
	*goa.Controller
	db *pg.DB
}

// NewDogController creates a dog controller.
func NewDogController(service *goa.Service, db *pg.DB) *DogController {
	return &DogController{Controller: service.NewController("DogController"), db: db}
}

// Create runs the create action.
func (c *DogController) Create(ctx *app.CreateDogContext) error {
	defer c.db.Conn().Close()
	id, err := database.GenerateID()
	if err != nil {
		return err
	}
	dog := database.Dog{DogID: id, DogName: ctx.Payload.Name, DogOwner: ctx.Payload.Owner}
	return c.db.Insert(&dog)
}

// Delete runs the delete action.
func (c *DogController) Delete(ctx *app.DeleteDogContext) error {
	defer c.db.Conn().Close()
	params := ctx.Params[database.DogIDField]
	dogID, err := database.UUID(params[0])
	if err != nil {
		return err
	}
	dog := database.Dog{DogID: dogID}
	return c.db.Delete(&dog)
}
