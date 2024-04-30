package main

import (
	"log"

	"github.com/seanrmurphy/cf-race-api/api"
	"github.com/seanrmurphy/cf-race-api/service"
	"github.com/syumai/workers"
	_ "github.com/syumai/workers/cloudflare/d1" // register driver
)

func main() {
	// sec := SecurityHandlerImplementation{
	// 	DatabaseEnvVarName: "DB",
	// }

	// Create service instance.
	s := &service.ServerImplementation{
		DatabaseEnvVarName: "DB",
	}

	// Create generated server.
	srv, err := api.NewServer(s)
	if err != nil {
		log.Fatal(err)
	}
	workers.Serve(srv)
}
