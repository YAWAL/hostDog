//go:generate goagen bootstrap -d github.com/YAWAL/hostDog/design

package main

import (
	"flag"
	"log"

	"github.com/YAWAL/hostDog/app"
	"github.com/YAWAL/hostDog/config"
	"github.com/YAWAL/hostDog/database"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Load configs
	configFile := flag.String("path", "config.json", "config file path")
	flag.Parse()
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("can't load config: %s", err)
	}

	// Create service
	service := goa.New(conf.ServiceName)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Create DB connection
	db := database.Postgres(conf.DB)

	// Mount "dog" controller
	c := NewDogController(service, db)
	app.MountDogController(service, c)
	// Mount "host" controller
	c2 := NewHostController(service, db)
	app.MountHostController(service, c2)

	// Start service
	if err := service.ListenAndServe(conf.Port); err != nil {
		service.LogError("startup", "err", err)
	}

}
