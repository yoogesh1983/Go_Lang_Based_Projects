package transport

import (
	"Profile/proto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func InitializeTransportLayer() *gin.Engine {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAccountServiceClient(conn)
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

	sub.GET("/subscribeLogin", func(ctx *gin.Context) {

		msg, err := client.SubscribeSignIn(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Println("SignInResponse: ", msg.Recv)
		}
	})

	return sub
}
