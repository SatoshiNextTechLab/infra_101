package v1

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type chatServiceServer struct {
	msg chan string
}

func NewChatServiceServer() ChatServiceServer {
	return &chatServiceServer{msg: make(chan string, 1000)}
}

func (s *chatServiceServer) Send(ctx context.Context, mssg *wrappers.StringValue) (*empty.Empty, error) {
	fmt.Printf("Sent %v", *mssg)
	s.msg <- mssg.Value
	return &empty.Empty{}, nil
}

func (s *chatServiceServer) Subscribe(e *empty.Empty, stream ChatService_SubscribeServer) error {
	fmt.Println("requested")
	for {
		m := <-s.msg
		n := Msg{Text: fmt.Sprintf(" hello %s", m)}
		log.Printf("Message sent: %+v", n)
		return nil
	}
}
