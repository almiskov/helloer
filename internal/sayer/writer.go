package sayer

import (
	"fmt"
	"io"

	"github.com/almiskov/helloer/internal/app"
)

var _ app.Sayer = &writerSayer{}

type writerSayer struct {
	w io.Writer
}

func NewWriterSayer(w io.Writer) *writerSayer {
	return &writerSayer{
		w: w,
	}
}

// Say implements app.Sayer.
func (s *writerSayer) Say(something string) {
	fmt.Fprintln(s.w, something)
}
