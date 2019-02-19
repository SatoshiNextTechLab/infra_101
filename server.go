package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/n3m0/v1"

	"google.golang.org/grpc"
)

func runServer(ctx context.Context, srv v1.ChatServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil
	}

	server := grpc.NewServer()
	v1.RegisterChatServiceServer(server, srv)
	return server.Serve(listen)
}

func main() {
	if err := runServer(context.Background(), v1.NewChatServiceServer(), "3000"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
