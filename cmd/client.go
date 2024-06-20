package main

import (
	"context"
	"fmt"
	"log"

	"github.com/prafitradimas/go-grpc-impl/internal/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port := 9000
	con, err := grpc.NewClient(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server on port: %d", port)
	}
	defer con.Close()

	client := chat.NewChatServiceClient(con)
	_, err = client.SendMessage(context.Background(), &chat.Message{Body: "Testing"})
	if err != nil {
		log.Fatalf("error sending message: %w", err)
	}
}
