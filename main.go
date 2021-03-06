package main

import (
	"log"

	"github.com/tomocy/archs/config"
	"github.com/tomocy/archs/infra/http/registerer"
	"github.com/tomocy/archs/infra/http/route"
	"github.com/tomocy/archs/infra/http/server"
	"github.com/tomocy/archs/registry"
)

func main() {
	config.Must(config.LoadConfig("./config.yml"))

	route.MapRoutes(config.Current.Host, config.Current.Port)
	registry := registry.NewHTTPRegistry()
	webRegi := registerer.NewWebRegisterer(registry.NewWebHandler())
	server := server.New(webRegi)
	if err := server.ListenAndServe(":" + config.Current.Port); err != nil {
		log.Fatalf("failed to listen and serve: %s\n", err)
	}
}
