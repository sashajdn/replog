package repositories

import (
	"context"

	"github.com/rotisserie/eris"

	"github.com/sashajdn/replog/pkg/application/replog"
	"github.com/sashajdn/replog/pkg/db"
	"github.com/sashajdn/replog/pkg/log"
)

var _ replog.ReplogRepository = &EntryRepository{}

func NewEntryRepository(logger *log.SugaredLogger, db *db.DB) *EntryRepository {
	return &EntryRepository{db: db}
}

type EntryRepository struct {
	logger *log.SugaredLogger
	db     *db.DB
}

func (r *EntryRepository) AppendLinesToEntry(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}

func (r *EntryRepository) CreateEntry(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}

func (r *EntryRepository) CreateEntryAndAppendLines(ctx context.Context, entry *replog.Entry) error {
	return eris.New("not implemented")
}
