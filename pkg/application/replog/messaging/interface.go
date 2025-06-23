package messaging

import (
	"context"
	"time"
)

type Message struct {
	ID           string    `json:"id"`
	ChannelID    string    `json:"channel_id"`
	SenderUserID string    `json:"sender_user_id"`
	Content      string    `json:"content"`
	Timestamp    time.Time `json:"timestamp"`
}

type ReceiverChannel chan *Message

type Receiver interface {
	Receive(ctx context.Context) (ReceiverChannel, error)
	Close() error
}

type Poster interface {
	SendPublicMessage(ctx context.Context, message *Message) error
	SendPrivateMessage(ctx context.Context, message *Message) error
}

type Client interface {
	Ping(ctx context.Context) error
	Receiver
	Poster
}
