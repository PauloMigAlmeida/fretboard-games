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

func StandardTuning() []*music.Note {
	eNote, _ := music.FindNote(music.E, music.Natural)
	bNote, _ := music.FindNote(music.B, music.Natural)
	gNote, _ := music.FindNote(music.G, music.Natural)
	dNote, _ := music.FindNote(music.D, music.Natural)
	aNote, _ := music.FindNote(music.A, music.Natural)

	return []*music.Note{
		eNote,
		bNote,
		gNote,
		dNote,
		aNote,
		eNote,
	}
}
