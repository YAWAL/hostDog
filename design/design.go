package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("hostDog", func() {
	Title("hosts and their dogs")
	Description("A simple goa service for hosts and their dogs")
	Scheme("http")
	BasePath("/hostdog")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("host", func() {
	BasePath("/hosts")
	Action("create", func() {
		Description("create host")
		Routing(POST(""))
		Payload(CreateHostPayload)
		Response(Created)
		Response(BadRequest)
	})
	Action("delete", func() {
		Description("delete host by id")
		Routing(DELETE("/:hostID"))
		Params(func() {
			Param("hostID", UUID, "Host ID")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest)
	})
	Action("getAll", func() {
		Description("get all hosts")
		Routing(GET(""))
		Response(OK, CollectionOf(HostsResponse))
		Response(NotFound)
	})
	Action("getDogsByHostID", func() {
		Description("get list of dogs by host's ID")
		Routing(GET("/:hostID/dogs"))
		Response(OK, func() {
			Status(200)
			Media(Dogs)
		})
		Response(NotFound)
		Response(BadRequest)
	})
	Action("getByID", func() {
		Description("get list of dogs by host's ID")
		Routing(GET("/:hostID"))
		Response(OK, func() {
			Status(200)
			Media(HostWithDogs)
		})
		Response(NotFound)
		Response(BadRequest)
	})
})

var _ = Resource("dog", func() {
	BasePath("/dogs")
	Action("create", func() {
		Description("create dog")
		Routing(POST(""))
		Payload(CreateDogPayload)
		Response(Created)
		Response(BadRequest)
	})
	Action("delete", func() {
		Description("Delete dog by id")
		Routing(DELETE("/:dogID"))
		Params(func() {
			Param("dogID", UUID, "Dog's ID")
		})
		Response(OK)
		Response(BadRequest)
		Response(NotFound)
	})
})

var HostsResponse = MediaType("hostsResponse", func() {
	Description("Host entity")
	Attributes(func() {
		Attribute("id", UUID, "host's id")
		Attribute("age", Integer, "host's age")
		Attribute("name", String, "host's name")
		Required("age", "name")
	})
	View("default", func() {
		Attribute("id", UUID, "host's id")
		Attribute("age", Integer, "host's age")
		Attribute("name", String, "host's name")
		Required("name")
	})
})

var DogMedia = MediaType("dogResponse", func() {
	Description("Dog entity")
	Attribute("id", UUID, "dog's id")
	Attribute("name", String, "dog's name")
	Attribute("owner", UUID, "id of the dog's host")
	View("default", func() {
		Attribute("id", UUID, "dog's id")
		Attribute("name", String, "dog's name")
		Attribute("owner", UUID, "id of the dog's host")
	})
})

var Dogs = MediaType("dogs", func() {
	Description("Dogs entity")
	Attribute("dogs", CollectionOf(DogMedia), "list of dogs")
	View("default", func() {
		Attribute("dogs", CollectionOf(DogMedia), "list of dogs")
	})
})

var HostWithDogs = MediaType("hostWithDogs", func() {
	Description("host and list of dogs")
	Attribute("id", UUID, "host's id")
	Attribute("hostAge", Integer, "host's age")
	Attribute("hostName", String, "host's name")
	Attribute("dogs", CollectionOf(DogMedia), "list of host's dogs")
	View("default", func() {
		Attribute("id", UUID, "host's id")
		Attribute("hostAge", Integer, "host's age")
		Attribute("hostName", String, "host's name")
		Attribute("dogs", Dogs, "list of host's dogs")
	})
})

var CreateHostPayload = Type("createHostPayload", func() {
	Description("payload for creation of host")
	Attribute("age", Integer, "host's age")
	Attribute("name", String, "host's name")
	Required("age", "name")
})

var CreateDogPayload = Type("createDogPayload", func() {
	Description("payload for creation of dog")
	Attribute("name", String, "dog's name")
	Attribute("owner", UUID, "host's ID")
	Required("name", "owner")
})
