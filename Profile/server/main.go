package main

import (
	"Profile/proto"
	service "Profile/server/service"
	"log"
	"net"
	"os"

	grpc "google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {

	// create a new gRPC server, use WithInsecure to allow http connections
	s := grpc.NewServer()

	// create an instance of the Service
	var grpcLog glog.LoggerV2 = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	var connections []*service.Connection

	srv := service.NewService(grpcLog, connections)

	// Register the service and server
	proto.RegisterBroadcastServer(s, srv)

	// You might want to disable into a production environment though!!
	reflection.Register(s)

	// Start to listen
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("Unable to create listener %v", err)
		os.Exit(1)
	}

	grpcLog.Info("Server running on port 4040")
	if e := s.Serve(listener); e != nil {
		log.Fatalf("Unable to create listener %v", e)
		os.Exit(1)
	}
}
