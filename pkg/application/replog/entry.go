package replog

import (
	"io"

	"github.com/rotisserie/eris"
)

type Entry struct {
	ID    string
	Lines []*Line
}

func (e *Entry) Append(line *Line) {
	e.Lines = append(e.Lines, line)
}

func (e *Entry) Flush(writer io.Writer) error {
	return eris.New("not implemented")
}

type Line struct {
	UserID  string
	Content string
}

func (e *Line) Flush(writer io.Writer) error {
	return eris.New("not implemented")
}
