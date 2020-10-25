package service

import (
	"Profile/proto"
	"Profile/server/data"
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
)

type server struct {
	log           hclog.Logger
	token         *data.JwtToken
	subscriptions map[proto.AccountService_SubscribeSignInServer][]*proto.SignInRequest
}

func NewService(l hclog.Logger, t *data.JwtToken) *server {
	s := &server{l, &data.JwtToken{Log: l, Token: ""}, make(map[proto.AccountService_SubscribeSignInServer][]*proto.SignInRequest)}
	go s.handleUpdates()
	return s
}

func (s *server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	s.log.Info("Handle request for SignIn", "Username", request.GetUsername(), "Password", request.GetPassword())

	r := &proto.SignInResponse{
		TransactionID: "57461347597412563612002484",
		Data: &proto.Data{
			JwtToken: "KXY58T325EWDH5ADF5A2DF8DFAFSDFS5DF5DFDF5DFG5SD5FG2F",
			Guid:     5,
		},
	}

	return r, nil
}

// SubscribeSignIn implments the gRPC bidirection streaming method for the server
func (s *server) SubscribeSignIn(source proto.AccountService_SubscribeSignInServer) error {

	// r := &proto.SignInResponse{
	// 	TransactionID: "57461347597412563612002484",
	// 	Data: &proto.Data{
	// 		JwtToken: "KXY58T325EWDH5ADF5A2DF8DFAFSDFS5DF5DFDF5DFG5SD5FG2F",
	// 		Guid:     10,
	// 	},
	// }

	// handle client messages
	for {
		rr, err := source.Recv() // Recv is a blocking method which returns on client data
		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			s.log.Info("Client has closed connection")
			return err
		}
		// any other error means the transport between the server and client is unavailable
		if err != nil {
			s.log.Error("Unale to read from client", "error", err)
			return err
		}

		s.log.Info("Handle client request", "request", rr)
		rrs, ok := s.subscriptions[source]
		if !ok {
			rrs = []*proto.SignInRequest{}
		}
		rrs = append(rrs, rr)
		s.subscriptions[source] = rrs
	}
	return nil
}

func (s *server) handleUpdates() {
	s.token.MonitorJWTStatus(30 * time.Second)
	s.log.Info("Got Updated rates")
}
