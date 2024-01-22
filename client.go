package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"google.golang.org/grpc"

	"chat-rpc/proto"
)

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)

	// Dial to the gRPC server (assuming it's running on localhost:8020)
	conn, err := grpc.Dial("localhost:8020", grpc.WithInsecure())
	if err != nil {
		l.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := proto.NewChatClient(conn)

	// Call the SendMessage RPC
	streamSend, err := client.SendMessage(context.Background())
	streamRecieve, err := client.RecieveMessage(context.Background())
	if err != nil {
		l.Fatalf("Error creating stream: %v", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		flag := &proto.MessageFlag{
			Flag: "1",
		}
		if err := streamRecieve.Send(flag); err != nil {
			l.Fatalf("Error sending message: %v", err)
		}
		for {
			// Receive and print responses from the server stream.
			res, err := streamRecieve.Recv()
			if err != nil {
				l.Printf(", Error receiving response: %v", err)
				break
			}
			message := res.GetText()
			if strings.ToUpper(message) == "CLOSE" {
				break
			}
			fmt.Printf("\nServer: %s\n", message)
		}
		streamRecieve.CloseSend()
		wg.Done()
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nMessage: ")
		scanner.Scan()
		messageText := scanner.Text()

		message := &proto.MessageSend{
			Text: messageText,
		}
		if strings.ToUpper(messageText) == "CLOSE" {
			if err := streamSend.Send(message); err != nil {
				l.Fatalf("Error sending message: %v", err)
			}
			break
		}

		if err := streamSend.Send(message); err != nil {
			l.Fatalf("Error sending message: %v", err)
		}
	}
	wg.Wait()
	// Read messages from the command prompt and send them

	// Close the stream
	streamSend.CloseSend()
}
