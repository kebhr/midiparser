package main

import "github.com/kebhr/midiparser"

func main() {
	midi, err := midiparser.New("../../midi1.mid")
	if err != nil {
		panic(err.Error())
	}
	if err := midi.Scan(); err != nil {
		panic(err)
	}
}
