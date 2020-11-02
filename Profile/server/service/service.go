package service

import (
	"Profile/proto"
	"context"
	"sync"

	glog "google.golang.org/grpc/grpclog"
)

type server struct {
	grpcLog    glog.LoggerV2
	threadPool []*NewConnection
}

func NewService(l glog.LoggerV2, c []*NewConnection) *server {
	s := &server{l, c}
	return s
}

type NewConnection struct {
	stream proto.AccountService_SubsribeForUpdatesServer
	id     string
	active bool
	error  chan error
}

//@Override
func (s *server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	s.grpcLog.Info("Handle request for SignIn", "Username", request.GetUsername(), "Password", request.GetPassword())

	r := &proto.SignInResponse{
		TransactionID: "57461347597412563612002484",
		Data: &proto.Data{
			JwtToken: "KXY58T325EWDH5ADF5A2DF8DFAFSDFS5DF5DFDF5DFG5SD5FG2F",
			Guid:     5,
		},
	}

	return r, nil
}

//@Override
func (s *server) SubsribeForUpdates(pconn *proto.Connection, stream proto.AccountService_SubsribeForUpdatesServer) error {
	conn := &NewConnection{
		stream: stream,
		id:     pconn.Client.Id,
		active: true,
		error:  make(chan error),
	}
	s.threadPool = append(s.threadPool, conn)
	return <-conn.error
}

//@Override
func (s *server) UpdateFirstName(ctx context.Context, msg *proto.Notification) (*proto.Close, error) {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, v := range s.threadPool {
		wg.Add(1)
		go func(notification *proto.Notification, conn *NewConnection) {
			defer wg.Done()

			if conn.active {
				err := conn.stream.Send(notification)
				s.grpcLog.Info("Updating a profile: ", conn.stream)

				if err != nil {
					s.grpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, v)
	}

	// since aboveis the go routine, we are makin here the another go routing to close the waitGroup and send a message to a channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// wait here until the ch return something here which means after the waitGroup is done. only after this channel get something from the pipe,
	// it let's the rest of the line to return
	<-ch
	return &proto.Close{}, nil
}
