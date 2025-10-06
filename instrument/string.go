package instrument

import "github.com/PauloMigAlmeida/fretboard-games/music"

type String struct {
	FretNotes []music.Note
}

func NewString(openNote *music.Note, numOfFrets int) *String {
	fretNotes := make([]music.Note, numOfFrets) // 0-indexed
	currNote := *openNote

	for i := range numOfFrets {
		fretNotes[i] = currNote
		nextNote, _ := currNote.NextHalfStepNote()
		currNote = *nextNote
	}

	return &String{
		FretNotes: fretNotes,
	}
}
