package main

import (
	transport "Profile/client/transport"
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func init() {
	wg = &sync.WaitGroup{}
}

func main() {

	// ***************** Create a connection ****************************
	client := transport.GetClient()
	user := transport.GetUser()
	err := transport.CreateConnection(client, user, wg)
	if err != nil {
		fmt.Errorf("error while trying to create a connection: %v", err)
	}

	// **************** Broadcast the Message *****************************
	err = transport.SendMessageToAll(client, user, wg)
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
