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

func InitializeTransportLayer() proto.ChatServiceClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewChatServiceClient(conn)
	return client
}

func StartChat(client proto.ChatServiceClient, wg *sync.WaitGroup) (string, error) {
	var streamerror error

	timestamp := time.Now()
	name := flag.String("N", "Kristy", "The name of the user")
	flag.Parse()
	id := sha256.Sum256([]byte(timestamp.String() + *name))
	encodedId := hex.EncodeToString(id[:])

	connection := &proto.Connection{
		User: &proto.User{
			Id:   encodedId,
			Name: *name,
		},
		Active: true,
	}

	stream, err := client.StartChat(context.Background(), connection)
	if err != nil {
		return "", fmt.Errorf("connection failed: %v", err)
	}

	wg.Add(1)
	go func(str proto.ChatService_StartChatClient) {
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

	fmt.Println("Successfully entered into a chatRoom with ID:", encodedId)
	fmt.Println("You are now eligible to Send and Receive all the new messages.")
	return encodedId, streamerror
}

func SendMessageToAll(client proto.ChatServiceClient, id string, wg *sync.WaitGroup) error {
	var broadCastError error

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := &proto.Message{
				Id:        id,
				Content:   scanner.Text(),
				Timestamp: time.Now().String(),
			}

			_, err := client.SendMessageToAll(context.Background(), msg)
			if err != nil {
				broadCastError = fmt.Errorf("connection failed: %v", err)
				break
			}
		}
	}()

	return broadCastError
}
