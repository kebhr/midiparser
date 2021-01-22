package midiparser

import (
	"os"
)

type file struct {
	file   *os.File
	cursor uint
}

func (*file) New(f *os.File) {
}
