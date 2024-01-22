package main

import (
	"bufio"
	"chat-rpc/proto"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type ChatServer struct {
	proto.UnimplementedChatServer
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func (cs *ChatServer) RecieveMessage(stream proto.Chat_RecieveMessageServer) error {
	l := log.New(os.Stdout, "", log.LstdFlags)

	req, err := stream.Recv()
	if req.Flag != "1" {
		fmt.Printf(req.Flag)
		fmt.Printf("Client has left!")
		return nil
	}
	if err == io.EOF {
		fmt.Printf("Client has left!")
		return nil
	}
	for {
		// Simulate server input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Server Message: ")
		scanner.Scan()
		serverMessage := scanner.Text()

		res := &proto.MessageRecieve{
			Text: serverMessage,
		}
		err := stream.Send(res)
		if err != nil {
			l.Fatalf("Error in Sending")
			return err
		}
	}
}

func (cs *ChatServer) SendMessage(stream proto.Chat_SendMessageServer) error {
	//l := log.New(os.Stdout, "", log.LstdFlags)

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("\nClient: %s\n", req.Text)
	}
}

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)
	grpc_server := grpc.NewServer()
	chat_server := NewChatServer()

	proto.RegisterChatServer(grpc_server, chat_server)

	lis, err := net.Listen("tcp", ":8020")
	if err != nil {
		l.Fatal("error in listening at 8020")
	}
	l.Println("gRPC server started on :8020")
	if err := grpc_server.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
	grpc_server.GracefulStop()
}
