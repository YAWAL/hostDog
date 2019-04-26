package main

import (
	"github.com/YAWAL/hostDog/app"
	"github.com/YAWAL/hostDog/database"
	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
)

// HostController implements the host resource.
type HostController struct {
	*goa.Controller
	db *pg.DB
}

// NewHostController creates a host controller.
func NewHostController(service *goa.Service, database *pg.DB) *HostController {
	return &HostController{Controller: service.NewController("HostController"), db: database}
}

// Create runs the create action.
func (c *HostController) Create(ctx *app.CreateHostContext) error {
	defer c.db.Conn().Close()
	id, err := database.GenerateID()
	if err != nil {
		return err
	}
	owner := database.Host{ID: id, Age: ctx.Payload.Age, Name: ctx.Payload.Name}
	return c.db.Insert(&owner)
}

// Delete runs the delete action.
func (c *HostController) Delete(ctx *app.DeleteHostContext) error {
	defer c.db.Conn().Close()
	params := ctx.Params[database.HostIDField]
	hostID, err := database.UUID(params[0])
	if err != nil {
		return err
	}
	host := database.Host{ID: hostID}
	return c.db.Delete(&host)
}

// GetAll runs the getAll action.
func (c *HostController) GetAll(ctx *app.GetAllHostContext) error {
	defer c.db.Conn().Close()
	var owners []database.Host
	res, err := c.db.Query(&owners, database.GetAllHost)
	if err != nil {
		return err
	}
	if res.RowsAffected() < 1 {
		return ctx.NotFound()
	}
	return ctx.OK(ConvertGetAllResult(owners))
}

// GetByID runs the getByID action.
func (c *HostController) GetByID(ctx *app.GetByIDHostContext) error {
	defer c.db.Conn().Close()
	var hostWithDogs []database.HostWithDogs
	params := ctx.Params[database.HostIDField]
	hostID, err := database.UUID(params[0])
	if err != nil {
		return err
	}
	res, err := c.db.Query(&hostWithDogs, database.GetHostByID, hostID)
	if err != nil {
		return err
	}
	if res.RowsAffected() < 1 {
		return ctx.NotFound()
	}
	return ctx.OK(ConvertGetByIDResult(hostWithDogs))
}

// GetDogsByHostID runs the getDogsByHostID action.
func (c *HostController) GetDogsByHostID(ctx *app.GetDogsByHostIDHostContext) error {
	defer c.db.Conn().Close()
	var dogs []database.Dog
	params := ctx.Params[database.HostIDField]
	res, err := c.db.Query(&dogs, database.GetDogsByHostID, params[0])
	if err != nil {
		return err
	}
	if res.RowsAffected() < 1 {
		return ctx.NotFound()
	}
	return ctx.OK(ConvertGetDogsByHostIDResult(dogs))
}

func ConvertGetAllResult(result []database.Host) (response []*app.Hostsresponse) {
	for _, res := range result {
		id := res.ID
		resp := app.Hostsresponse{ID: &id, Name: res.Name, Age: res.Age}
		response = append(response, &resp)
	}
	return response
}

func ConvertGetDogsByHostIDResult(result []database.Dog) *app.Dogs {
	var dogs app.DogresponseCollection
	response := app.Dogs{}
	for _, res := range result {
		id := res.DogID
		name := res.DogName
		owner := res.DogOwner
		resp := app.Dogresponse{ID: &id, Name: &name, Owner: &owner}
		dogs = append(dogs, &resp)
	}
	response.Dogs = dogs
	return &response
}

func ConvertGetByIDResult(result []database.HostWithDogs) *app.Hostwithdogs {
	var dogs app.DogresponseCollection
	response := app.Hostwithdogs{ID: &result[0].ID, HostName: &result[0].Name, HostAge: &result[0].Age}
	for _, res := range result {
		id := res.DogID
		name := res.DogName
		resp := app.Dogresponse{ID: &id, Name: &name}
		dogs = append(dogs, &resp)
	}
	response.Dogs = dogs
	return &response
}
