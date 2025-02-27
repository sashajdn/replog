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

	discordClient := aValidDiscordClient()

	err := discordClient.Ping(ctx)
	require.NoError(t, err)

	someMessage := someValidMessage()
	err = discordClient.SendPublicMessage(ctx, someMessage)
	require.NoError(t, err)
}

func TestDiscordClientSendPrivateMessage(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	discordClient := aValidDiscordClient()

	err := discordClient.Ping(ctx)
	require.NoError(t, err)

	someMessage := someValidMessage()
	err = discordClient.SendPrivateMessage(ctx, someMessage)
	require.NoError(t, err)
}

func aValidDiscordClient() *discord.DiscordClient {
	return discord.NewDiscordClient()
}

func someValidMessage() *messaging.Message {
	return &messaging.Message{}
}
