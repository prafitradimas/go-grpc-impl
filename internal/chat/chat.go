package chat

import (
	"context"
	"log"
)

type server struct{}

func NewServer() ChatServiceServer {
	return &server{}
}

func (srv *server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s\n", message.GetBody())
	return nil, nil
}

func (srv *server) mustEmbedUnimplementedChatServiceServer() {}
