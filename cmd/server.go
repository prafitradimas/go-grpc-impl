package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prafitradimas/go-grpc-impl/internal/chat"
	"github.com/prafitradimas/go-grpc-impl/pkg/assert"
	"google.golang.org/grpc"
)

func main() {
	port := 9000

	tcp, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	assert.NoError(err, fmt.Sprintf("Failed to listen on port: %d", port))

	grpcServer := grpc.NewServer()
	chatServer := chat.NewServer()
	chat.RegisterChatServiceServer(grpcServer, chatServer)

	fmt.Printf("Serving gRPC server on port: %d\n", port)
	err = grpcServer.Serve(tcp)
	if err != nil {
		log.Fatalf("could not serve gRPC server on port: %d", port)
	}
}
