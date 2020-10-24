package main

import (
	"Profile/client/transport"
	"fmt"
	"log"
)

func main() {
	client := transport.InitializeTransportLayer()

	fmt.Println("Client Running on port 8080")

	if err := client.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
