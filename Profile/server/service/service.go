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
	return &server{
		log: l,
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	s.log.Info("Handle request for Add", "FirstValue", request.GetA(), "SecondValue", request.GetB())
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	s.log.Info("Handle request for Multiply", "FirstValue", request.GetA(), "SecondValue", request.GetB())
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
