package main

import (
	"Profile/proto"
	"Profile/server/service"
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// create a new gRPC server, use WithInsecure to allow http connections
	s := grpc.NewServer()

	// create an instance of the Service
	log := hclog.Default()
	srv := service.NewService(log)

	// Register the service and server
	proto.RegisterAddServiceServer(s, srv)

	// You might want to disable into a production environment though!!
	reflection.Register(s)

	// Start to listen
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	fmt.Println("Server running on port 4040")
	if e := s.Serve(listener); e != nil {
		log.Error("Unable to serve listener", "error", e)
		os.Exit(1)
	}

}
