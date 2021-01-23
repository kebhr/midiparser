package midiparser

import (
	"os"
)

type file struct {
	file   *os.File
	cursor int
}

func (f *file) Open(of *os.File) *file {
	return &file{
		file:   of,
		cursor: 0,
	}
}

func (f *file) read(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := f.file.Read(buf); err != nil {
		return nil, err
	}
	f.cursor += n
	return buf, nil
}
