package replog

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/rotisserie/eris"
)

type (
	UserID    string
	ChannelID string
	MessageID string
)

type ReplogRepository interface {
	AppendLinesToEntry(ctx context.Context, entry *Entry) error
	CreateEntry(ctx context.Context, entry *Entry) error
	CreateEntryAndAppendLines(ctx context.Context, entry *Entry) error
	CreateChannel(ctx context.Context, channelName string, messagingProviderType MessagingProviderType) error
	ReadEntryFromUserAndChannel(ctx context.Context, userID string)
}

type UserRepository interface {
	CreateUser(ctx context.Context, username string, messagingProvider MessagingProvider) error
}

type MessagingProviderRepository interface {
	UpsertLastSeenMessage(ctx context.Context, messageID MessageID, userID UserID, channelID ChannelID) error
	ReadLastSeenMessage(ctx context.Context, userID UserID, channelID ChannelID) error
	ReadAllLastSeeNMessagesByUser(ctx context.Context, userID UserID) error
}

type User struct {
	ID                  string                `db:"id"`
	Username            string                `db:"username"`
	MessagingProviderID string                `db:"messaging_provider_id"`
	MessagingProvider   MessagingProviderType `db:"messaging_provider"`
	CreatedAt           time.Time             `db:"created_at"`
	UpdatedAt           time.Time             `db:"updated_at"`
}

func NewMessagingProvider(t MessagingProviderType, id string) MessagingProvider {
	return MessagingProvider{
		Type: t,
		ID:   id,
	}
}

type MessagingProvider struct {
	Type MessagingProviderType
	ID   string
}

var _ driver.Valuer = (*MessagingProviderType)(nil)
var _ sql.Scanner = (*MessagingProviderType)(nil)

type MessagingProviderType string

const (
	MessagingProviderDiscord MessagingProviderType = `discord`
)

func (m *MessagingProviderType) Scan(value any) error {
	if value == nil {
		*m = ""
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return eris.New(`scan MessagingProvider`)
	}

	*m = MessagingProviderType(string(b))
	return nil
}

func (m MessagingProviderType) Value() (driver.Value, error) {
	return string(m), nil
}
