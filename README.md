# gRPC Chat Application

This repository contains a simple two-way chat application implemented in Golang using gRPC. The chat application consists of a server and a client, allowing users to exchange messages in real-time.

## Features

- Bidirectional communication: Clients can send messages to the server, and the server can respond with messages.
- Interactive user input: Both the client and server accept user input from the command prompt.
- Graceful termination: Users can close the chat session by typing "CLOSE" in the client.

## Directory Structure

- `client/`: Contains the Golang code for the chat client.
- `server/`: Contains the Golang code for the chat server.
- `proto/`: Contains the protocol buffer (.proto) file defining the gRPC service.

## Usage

1. Clone the repository: `git clone <repository-url>`
2. Run the server: Navigate to the `server/` directory and run `go run server.go`.
3. Run the client: Navigate to the `client/` directory and run `go run client.go`.

Feel free to customize and expand the code for additional features or use it as a basis for more complex chat applications.

## Note

- This is a basic implementation for educational purposes. For a production environment, consider adding security features, error handling, and optimizations.
