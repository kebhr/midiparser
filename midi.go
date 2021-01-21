package midiparser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// Midi represents SMF Format 1
type Midi struct {
	file    *os.File
	timeSig *TimeSignature
}

// TimeSignature represents beat pattern
type TimeSignature struct {
	Numerator   uint
	Denominator uint
}

// New return a new struct
func New(fp string) (*Midi, error) {
	f, err := os.Open("midi1.mid")
	if err != nil {
		return &Midi{}, err
	}

	m := Midi{}
	m.file = f

	return &m, nil
}

// Scan scans midi file
func (m *Midi) Scan() error {
	return nil
}

