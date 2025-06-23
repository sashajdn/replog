package replog

import (
	"io"
	"time"

	"github.com/rotisserie/eris"
)

type Entry struct {
	ID        EntryID
	UserID    UserID
	Content   string
	Parent    *Entry
	Children  []*Entry
	CreatedAt time.Time
	UpdatedAt time.Time
	// TODO: add sub entry
	// TODO add tags & other metadata such as a headline
}

func (e *Entry) Flush(writer io.Writer) error {
	return eris.New("not implemented")
}
