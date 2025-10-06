package instrument

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/music"
)

type Fretboard struct {
	Strings []*String
}

func NewFretboard(numOfFrets int, tuning []*music.Note) *Fretboard {
	strings := make([]*String, len(tuning))

	for i := range len(strings) {
		strings[i] = NewString(tuning[i], numOfFrets)
	}

	return &Fretboard{
		Strings: strings,
	}
}

func (f *Fretboard) GetNoteAt(stringNumber int, fretNumber int) (*music.Note, error) {
	if stringNumber < 1 {
		return nil, fmt.Errorf("strings are 1-indexed to be relation to how we number them in real world")
	}

	if stringNumber > len(f.Strings) {
		return nil, fmt.Errorf("string '%d' doesn't exist", stringNumber)
	}

	if fretNumber < 0 || fretNumber > len(f.Strings[0].FretNotes)-1 {
		return nil, fmt.Errorf("fret '%d' doesn't exist", stringNumber)
	}

	// strings are 1-indexed
	return &f.Strings[stringNumber-1].FretNotes[fretNumber], nil
}
