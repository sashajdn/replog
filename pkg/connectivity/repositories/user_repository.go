package repositories

import (
	"context"

	"github.com/rotisserie/eris"
	"github.com/sashajdn/replog/pkg/application/replog"
	"github.com/sashajdn/replog/pkg/db"
	"github.com/sashajdn/replog/pkg/log"
)

var _ replog.UserRepository = &UserRepository{}

func NewUserRepository(logger *log.SugaredLogger, db *db.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	logger *log.SugaredLogger
	db     *db.DB
}

func (u *UserRepository) CreateUser(
	ctx context.Context,
	username,
	messagingProvider replog.MessagingProvider,
) error {
	const query = `
    WITH ins AS (
        INSERT INTO "user" (username, messaging_provider_id, messaging_provider)
        VALUES ($1, $2, $3)
        ON CONFLICT (messaging_provider_id) DO NOTHING
        RETURNING id, username, messaging_provider_id, messaging_provider, created_at, updated_at
    )
    SELECT
        id,
        username,
        messaging_provider_id,
        messaging_provider,
        created_at,
        updated_at
    FROM ins
    UNION ALL
    SELECT
        id,
        username,
        messaging_provider_id,
        messaging_provider,
        created_at,
        updated_at
    FROM "user"
    WHERE messaging_provider_id = $2
    LIMIT 1;
    `

	var user replog.User
	if err := u.db.QueryRow(
		ctx,
		query,
		username,
		messagingProvider.ID,
		messagingProvider.Type,
	).Scan(
		&user.ID,
		&user.Username,
		&user.MessagingProviderID,
		&user.MessagingProvider,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return eris.Wrap(err, "failed to create user")
	}

	return nil
}
