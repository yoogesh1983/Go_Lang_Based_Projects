package main

import (
	transport "Profile/client/transport"
	"fmt"
	"log"
	"sync"
)

var wg *sync.WaitGroup

func init() {
	wg = &sync.WaitGroup{}
}

func main() {

	// ***************** Create a client ****************************
	client := transport.InitializeTransportLayer()

	//***************** Create Restful api **************************
	go func() {
		restClient := transport.StartRestAPI(client)
		fmt.Println("Client Running on port 8080")
		if err := restClient.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// *********************** Start Subscription ****************************
	id, err := transport.SubsribeForUpdates(client, wg)
	if err != nil {
		fmt.Errorf("error while trying to create a connection: %v", err)
	}

	// **************** Broadcast the Message *****************************
	err = transport.UpdateFirstName(client, id, wg)
	if err != nil {
		fmt.Errorf("error while broadcasting message: %v", err)
	}

	// *************** Wait until the channel gets a message **************
	ch := make(chan int)
	// since above are the go routine, we are making here the another go routing to close the waitGroup and send a message to a channel
	go func() {
		wg.Wait()
		close(ch)
	}()
	// wait here until the ch return something here which means after the waitGroup is done. only after this channel get something from the pipe,
	// it let's the rest of the line to return
	<-ch
}
