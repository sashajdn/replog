package messaging

import "context"

type Message struct {
	ID           string `json:"id"`
	ChannelID    string `json:"channel_id"`
	SenderUserID string `json:"sender_user_id"`
	Content      string `json:"content"`
}

type Poster interface {
	SendPublicMessage(ctx context.Context, message *Message) error
	SendPrivateMessage(ctx context.Context, message *Message) error
}

type Client interface {
	Ping(ctx context.Context) error
	Poster
}
