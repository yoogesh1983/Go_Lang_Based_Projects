package main

import (
	dao "Profile/server/Database"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	//Start server
	initServer()
}

func initServer() {
	var api = new(dao.API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, Err := net.Listen("tcp", ":4040")

	if Err != nil {
		log.Fatal("Listener Error", Err)
	}
	log.Printf("Serving RPC on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving", err)
	}
}
