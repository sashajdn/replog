package discord

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"

	"github.com/sashajdn/replog/pkg/application/replog/messaging"
	"github.com/sashajdn/replog/pkg/log"
)

var _ messaging.Client = &DiscordClient{}

const defaultReceiverChannelSize = 100_000

type DiscordClient struct {
	session      *discordgo.Session
	receiverSize int
	receiver     messaging.ReceiverChannel
	cfg          ClientConfig
	logger       *log.SugaredLogger
}

type ClientConfig struct {
	Token               string
	ReceiverChannelSize int
	Logger              *log.SugaredLogger
}

type DiscordClientOption func(client *DiscordClient)

func NewClient(cfg ClientConfig, clientOpts ...DiscordClientOption) (*DiscordClient, error) {
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, eris.Wrap(err, `create discord session`)
	}

	var receiverSize = cfg.ReceiverChannelSize
	if receiverSize == 0 {
		receiverSize = defaultReceiverChannelSize
	}

	return &DiscordClient{
		logger:       cfg.Logger,
		session:      session,
		cfg:          cfg,
		receiverSize: receiverSize,
	}, nil
}

func (d *DiscordClient) Receive(ctx context.Context) (messaging.ReceiverChannel, error) {
	if d.session == nil {
		return nil, eris.New("session is nil")
	}

	d.logger.Info("Opening discord client session...")
	if err := d.session.Open(); err != nil {
		return nil, eris.Wrap(err, "session open")
	}

	d.receiver = make(messaging.ReceiverChannel, d.receiverSize)
	d.session.AddHandler(d.handleMessageCreate)

	return d.receiver, nil
}

func (d *DiscordClient) Close() error {
	if d.session == nil {
		return nil
	}

	return d.session.Close()
}

func (d *DiscordClient) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	d.logger.With(zap.String("message_id", m.ID)).Debugf("Received message create")
	select {
	case d.receiver <- &messaging.Message{
		ID:           m.ID,
		Content:      m.Content,
		ChannelID:    m.ChannelID,
		SenderUserID: m.Author.ID,
		Timestamp:    m.Timestamp,
	}:
	default:
		d.logger.Warn("discord client receiver full")
	}
}

func (d *DiscordClient) Ping(ctx context.Context) error {
	return eris.New("not implemented")
}

func (d *DiscordClient) SendPublicMessage(ctx context.Context, message *messaging.Message) error {
	if _, err := d.session.ChannelMessageSend(
		message.ChannelID,
		message.Content,
		discordgo.WithContext(ctx),
	); err != nil {
		return eris.Wrap(err, "channel send message")
	}

	return nil
}

func (d *DiscordClient) SendPrivateMessage(ctx context.Context, message *messaging.Message) error {
	channel, err := d.session.UserChannelCreate(message.SenderUserID)
	if err != nil {
		return fmt.Errorf("user channel create")
	}

	if _, err := d.session.ChannelMessageSend(
		channel.ID,
		message.Content,
		discordgo.WithContext(ctx),
	); err != nil {
		err := eris.Wrap(err, "channel send message")
		if err != nil {
			return err
		}
	}

	return nil
}
