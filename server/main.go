package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/n3m0/chat/api"
	"google.golang.org/grpc"
)

type Server struct {
	//
}

func (s Server) ChatStream(srv api.Chat_ChatStreamServer) error {

	log.Println("start new server")
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}

		if err != nil {
			log.Println("recieve error")
			continue
		}

		fmt.Println("recieved message %v", req.Message)
		resp := api.Note{Message: "msg from server"}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("sent msg")
	}
}

func main() {
	a, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("%v", err)
	}

	s := grpc.NewServer()
	api.RegisterChatServer(s, Server{})

	if err := s.Serve(a); err != nil {
		log.Fatalf("%s", err)
	}
}
