package transport

import (
	"Profile/proto"
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func InitializeTransportLayer() proto.AccountServiceClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAccountServiceClient(conn)
	return client
}

func StartRestAPI(client proto.AccountServiceClient) *gin.Engine {
	sub := gin.Default()

	sub.GET("/login", func(ctx *gin.Context) {
		signInRequest := &proto.SignInRequest{Username: "admin@gmail.com", Password: "1234"}
		if response, err := client.SignIn(ctx, signInRequest); err == nil {
			fmt.Println("SignInResponse: ", response)
			ctx.JSON(http.StatusOK, gin.H{
				"transactionId": response.TransactionID,
				"data": gin.H{
					"jwtToken": response.Data.JwtToken,
					"guid":     response.Data.Guid,
				},
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	return sub
}

func SubsribeForUpdates(client proto.AccountServiceClient, wg *sync.WaitGroup) (string, error) {
	var streamerror error

	timestamp := time.Now()
	name := flag.String("N", "totalwine", "This will be the default name of the client")
	flag.Parse()
	id := sha256.Sum256([]byte(timestamp.String() + *name))
	encodedID := hex.EncodeToString(id[:])

	connection := &proto.Connection{
		Client: &proto.Client{
			Id:   encodedID,
			Name: *name,
		},
		Active: true,
	}

	stream, err := client.SubsribeForUpdates(context.Background(), connection)
	if err != nil {
		return "", fmt.Errorf("connection failed: %v", err)
	}

	wg.Add(1)
	go func(str proto.AccountService_SubsribeForUpdatesClient) {
		defer wg.Done()

		for {
			notification, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("Error reading message: %v", err)
				break
			}

			if encodedID == notification.Id {
				fmt.Printf("You : %s\n", notification.Profile)
			} else {
				fmt.Printf("%v : %s\n", notification.Id, notification.Profile)
			}
		}
	}(stream)

	fmt.Println("Successfully subscribed for updateProfile with clientID:", encodedID)
	fmt.Println("You are now eligible for update or receive any update profile information")
	return encodedID, streamerror
}

func UpdateFirstName(client proto.AccountServiceClient, id string, wg *sync.WaitGroup) error {
	var broadCastError error

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			notification := &proto.Notification{
				Id:        id,
				Timestamp: time.Now().String(),
				Profile: &proto.Profile{
					FirstName: scanner.Text(),
					LastName:  "updated-lastName",
				},
			}

			_, err := client.UpdateFirstName(context.Background(), notification)
			if err != nil {
				broadCastError = fmt.Errorf("connection failed: %v", err)
				break
			}
		}
	}()

	return broadCastError
}
