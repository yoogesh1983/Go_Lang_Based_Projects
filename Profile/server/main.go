package main

import (
	"Profile/proto"
	"Profile/server/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &service.Server{})
	reflection.Register(srv)

	fmt.Println("Server running on port 4040")

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}

}
