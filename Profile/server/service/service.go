package service

import (
	"Profile/proto"
	"context"
	"sync"

	glog "google.golang.org/grpc/grpclog"
)

type server struct {
	grpcLog    glog.LoggerV2
	threadPool []*Connection
}

type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

func NewService(l glog.LoggerV2, c []*Connection) *server {
	s := &server{l, c}
	return s
}

func (s *server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}
	s.threadPool = append(s.threadPool, conn)
	return <-conn.error
}

func (s *server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for _, v := range s.threadPool {
		wg.Add(1)
		go func(msg *proto.Message, conn *Connection) {
			defer wg.Done()

			if conn.active {
				err := conn.stream.Send(msg)
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
