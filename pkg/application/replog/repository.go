package replog

import (
	"context"
)

type Repository interface {
	AppendLinesToEntry(ctx context.Context, entry *Entry) error
	CreateEntry(ctx context.Context, entry *Entry) error
	CreateEntryAndAppendLines(ctx context.Context, entry *Entry) error
}
