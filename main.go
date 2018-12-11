package main

import (
	"github.com/go-land/job-service/handlers"
	"github.com/go-land/job-service/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("job"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	job.RegisterJobServiceHandler(srv.Server(), handlers.NewJobHandler())

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Can't properly start server: %v\n", err)
	}
}
