//go:build integration
// +build integration

package discord_test

import (
	"context"
	"testing"
	"time"

	"github.com/sashajdn/replog/pkg/application/replog/messaging"
	"github.com/sashajdn/replog/pkg/connectivity/discord"
	"github.com/stretchr/testify/require"
)

func TestDiscordClientSendPublicMessage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	discordClient := aValidDiscordClient(t)

	someMessage := someValidMessage()
	err := discordClient.SendPublicMessage(ctx, someMessage)
	require.NoError(t, err)
}

func TestDiscordClientSendPrivateMessage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	discordClient := aValidDiscordClient(t)

	someMessage := someValidMessage()
	err := discordClient.SendPrivateMessage(ctx, someMessage)
	require.NoError(t, err)
}

func aValidDiscordClient(t *testing.T) *discord.DiscordClient {
	cfg := discord.ClientConfig{}
	client, err := discord.NewClient(cfg)
	require.NoError(t, err, `failed to create a valid discord client`)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx)
	require.NoError(t, err, `failed to ping a valid discord client`)

	return client
}

func someValidMessage() *messaging.Message {
	return &messaging.Message{}
}
