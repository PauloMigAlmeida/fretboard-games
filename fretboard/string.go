package fretboard

type String struct {
	FretNotes []Note
}

func NewString(openNote *Note, numOfFrets int) *String {
	fretNotes := make([]Note, numOfFrets+1) // 0-indexed
	currNote := *openNote

	for i := range numOfFrets + 1 {
		fretNotes[i] = currNote
		nextNote, _ := currNote.NextHalfStepNote()
		currNote = *nextNote
	}

	return &String{
		FretNotes: fretNotes,
	}
}
