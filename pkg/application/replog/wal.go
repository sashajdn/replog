package replog

import (
	"fmt"
	"io"

	"github.com/rotisserie/eris"
)

var _ io.Writer = &WAL{}

const defaultWALBufferSize = 100

func NewWAL() *WAL {
	return &WAL{
		entry: nil,
	}
}

type WAL struct {
	entry *Entry
}

func (w *WAL) CreateNewEntry(id string) error {
	if w.entry != nil {
		return fmt.Errorf("entry already exists")
	}

	w.entry = &Entry{
		Lines: make([]*Line, defaultWALBufferSize),
		ID:    id,
	}

	return nil
}

func (w *WAL) Append(line *Line) error {
	if err := line.Flush(w); err != nil {
		return eris.Wrap(err, "failed to flush line")
	}

	w.entry.Lines = append(w.entry.Lines, line)

	if len(w.entry.Lines) >= defaultWALBufferSize {
		if err := w.entry.Flush(w); err != nil {
		}
	}

	return nil
}

func (w *WAL) Reset() {
	w.entry = nil
}

func (w *WAL) Write(p []byte) (n int, err error)
