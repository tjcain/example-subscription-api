package main

import (
	"log"
	"net/http"

	"github.com/tjcain/example-subscription-api/pkg/adding"
	"github.com/tjcain/example-subscription-api/pkg/config"
	"github.com/tjcain/example-subscription-api/pkg/http/rest"
	"github.com/tjcain/example-subscription-api/pkg/storage"
)

func main() {
	// set up dependencies:
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("could not get config: %s", err)
	}

	s, err := storage.NewPostgres(cfg.Database)
	if err != nil {
		log.Fatalf("could not get reach postgres: %s", err)
	}

	// set up service.
	a := adding.NewService(s)
	// ... add other services

	// setup router
	router := rest.Handler(a) // update to add new services.
	if err := http.ListenAndServe(cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
