package transport

import (
	"Profile/proto"
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"sync"

	"encoding/hex"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func GetClient() proto.BroadcastClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewBroadcastClient(conn)
	return client
}

func GetUser() *proto.User {
	timestamp := time.Now()
	name := flag.String("N", "Kristy", "The name of the user")
	flag.Parse()
	id := sha256.Sum256([]byte(timestamp.String() + *name))

	user := proto.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}
	return &user
}

func CreateConnection(client proto.BroadcastClient, user *proto.User, wg *sync.WaitGroup) error {
	var streamerror error

	stream, err := client.CreateStream(context.Background(), &proto.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wg.Add(1)
	go func(str proto.Broadcast_CreateStreamClient) {
		defer wg.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("Error reading message: %v", err)
				break
			}
			fmt.Printf("%v : %s\n", msg.Id, msg.Content)
		}
	}(stream)

	return streamerror
}

func SendMessage(client proto.BroadcastClient, user *proto.User, wg *sync.WaitGroup) error {
	var broadCastError error

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := &proto.Message{
				Id:        user.Id,
				Content:   scanner.Text(),
				Timestamp: time.Now().String(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				broadCastError = fmt.Errorf("connection failed: %v", err)
				break
			}
		}
	}()

	return broadCastError
}
