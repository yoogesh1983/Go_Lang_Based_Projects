package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handlers "Products/src/handlers"

	"github.com/nicholasjackson/env"
)

func main() {

	httpAddr := env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create your own product handler and register that handlers to the ServeMux() handler which is the default one provided by golang. later you will replace it be Mux Router
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := http.Server{
		Addr:         *httpAddr,         // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		//err := http.ListenAndServe(*httpAddr, sm)
		err := s.ListenAndServe() // This way you can define a timeout also ...
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
