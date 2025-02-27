package repositories

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/sashajdn/replog/pkg/application/replog"
	"github.com/sashajdn/replog/pkg/db"
)

var _ replog.Repository = &ReplogRepository{}

func NewReplogRepository(db *db.DB) *ReplogRepository {
	return &ReplogRepository{db: db}
}

type ReplogRepository struct {
	db *db.DB
}

func (r *ReplogRepository) AppendLinesToEntry(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}

func (r *ReplogRepository) CreateEntry(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}

func (r *ReplogRepository) CreateEntryAndAppendLines(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}
