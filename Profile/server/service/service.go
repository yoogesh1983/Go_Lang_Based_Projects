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
	stream proto.ChatService_StartChatServer
	id     string
	active bool
	error  chan error
}

//@Override
func (s *server) StartChat(pconn *proto.Connection, stream proto.ChatService_StartChatServer) error {
	conn := &NewConnection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}
	s.threadPool = append(s.threadPool, conn)
	return <-conn.error
}

//@Override
func (s *server) SendMessageToAll(ctx context.Context, msg *proto.Notification) (*proto.Close, error) {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, v := range s.threadPool {
		wg.Add(1)
		go func(notification *proto.Notification, conn *NewConnection) {
			defer wg.Done()

			if conn.active {
				err := conn.stream.Send(notification)
				s.grpcLog.Info("Sending message to: ", conn.stream)

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
