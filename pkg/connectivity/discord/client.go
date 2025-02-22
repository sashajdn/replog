package discord

import (
	"context"
	"fmt"

	"github.com/sashajdn/replog/pkg/application/replog/messaging"
)

var _ messaging.Client = &DiscordClient{}

type DiscordClient struct{}

type DiscordClientOption func(client *DiscordClient)

func NewDiscordClient(clientOpts ...DiscordClientOption) *DiscordClient {
	return &DiscordClient{}
}

func (d *DiscordClient) Ping(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}

func (d *DiscordClient) SendPublicMessage(ctx context.Context, message *messaging.Message) error {
	return fmt.Errorf("not implemented")
}

func (d *DiscordClient) SendPrivateMessage(ctx context.Context, message *messaging.Message) error {
	return fmt.Errorf("not implemented")
}
