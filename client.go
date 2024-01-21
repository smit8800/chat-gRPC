package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		l.Fatalf("Error creating stream: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Message: ")
		scanner.Scan()
		messageText := scanner.Text()

		message := &proto.MessageSend{
			Text: messageText,
		}
		if strings.ToUpper(messageText) == "CLOSE" {
			break
		}

		if err := stream.Send(message); err != nil {
			l.Fatalf("Error sending message: %v", err)
		}

		// Receive and print responses from the server stream.
		res, err := stream.Recv()
		if err != nil {
			l.Printf(", Error receiving response: %v", err)
			break
		}
		fmt.Printf("Server: %s\n", res.GetText())
	}

	// Read messages from the command prompt and send them

	// Close the stream
	stream.CloseSend()
}
