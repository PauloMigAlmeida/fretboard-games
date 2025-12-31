package instrument

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"slices"
	"strings"
)

type Fretboard struct {
	Strings []*String
}

func NewFretboard(numOfFrets int, tuning []*music.Note) *Fretboard {
	fretStrings := make([]*String, len(tuning))

	for i := range len(fretStrings) {
		fretStrings[i] = NewString(tuning[i], numOfFrets)
	}

	return &Fretboard{
		Strings: fretStrings,
	}
}

func (f *Fretboard) GetNoteAt(stringNumber int, fretNumber int) (*music.Note, error) {
	if stringNumber < 1 {
		return nil, fmt.Errorf("strings are 1-indexed to be like how we number them in real world")
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

func (f *Fretboard) DrawFretboard(notes []*music.Note, ignoreStrings []int) (string, error) {
	var sb strings.Builder

	if len(f.Strings) == 0 || len(f.Strings[0].FretNotes) == 0 {
		return "", fmt.Errorf("fretboard has no strings or frets")
	}

	// Header
	sb.WriteString("|")
	for idx := range f.Strings[0].FretNotes {
		sb.WriteString(fmt.Sprintf(" %-3d|", idx))
	}
	sb.WriteString("\n")

	// Body
	for strIdx, strEl := range f.Strings {
		sb.WriteString("|")
		for _, note := range strEl.FretNotes {

			if slices.Contains(ignoreStrings, strIdx+1) {
				sb.WriteString(fmt.Sprintf(" %-3s", "-"))
			} else {
				found := false
				for _, n := range notes {
					if n.Equals(&note) {
						sb.WriteString(fmt.Sprintf(" %-3s", "X"))
						found = true
					}
				}

				if !found {
					sb.WriteString(fmt.Sprintf(" %-3s", "-"))
				}
			}

			sb.WriteString("|")
		}

		if strIdx < len(f.Strings)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
