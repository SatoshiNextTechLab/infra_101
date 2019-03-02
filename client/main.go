package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/n3m0/chat/api"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := api.NewChatClient(conn)
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("open stream err %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		in, _ := reader.ReadString('\n')

		req := api.Note{Message: in}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("can not send %v", err)
		}

		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not revieve %v", err)
			}
			mess := resp.Message
			log.Printf("message recieved %v", mess)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("done")
}
