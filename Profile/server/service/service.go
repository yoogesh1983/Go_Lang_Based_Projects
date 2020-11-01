package service

import (
	"Profile/proto"
	"context"

	"github.com/hashicorp/go-hclog"
)

type server struct {
	log hclog.Logger
}

func NewService(l hclog.Logger) *server {
	s := &server{l}
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
