package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"io"
	"math/rand"
	"strings"
)

type FindNoteGame struct {
	Fretboard *instrument.Fretboard
	// game variables
	NotesAmount   int
	StringsAmount int
	// OS stuff
	StdIn io.Reader
}

func NewFindNoteGame(fretboard *instrument.Fretboard, stdIn io.Reader) *FindNoteGame {
	return &FindNoteGame{
		Fretboard:     fretboard,
		NotesAmount:   0,
		StringsAmount: 0,
		StdIn:         stdIn,
	}
}

func (f *FindNoteGame) Configure() error {
	var notesAmount int
	fmt.Print("How many notes do you want to find?: ")
	_, err := fmt.Fscanf(f.StdIn, "%d\n", &notesAmount)
	if err != nil {
		return err
	}
	f.NotesAmount = notesAmount

	var stringsAmount int
	fmt.Print("Across how many strings do you want to find notes?: ")
	_, err = fmt.Fscanf(f.StdIn, "%d\n", &stringsAmount)
	if err != nil {
		return err
	}
	f.StringsAmount = stringsAmount

	if f.StringsAmount < 1 || f.StringsAmount > len(f.Fretboard.Strings) {
		return fmt.Errorf("invalid strings amount, has to be between 1 and %d", len(f.Fretboard.Strings))
	}

	if f.NotesAmount < 1 || f.NotesAmount > len(f.Fretboard.Strings[0].FretNotes) {
		return fmt.Errorf("invalid notes amount, has to be between 1 and %d", len(f.Fretboard.Strings[0].FretNotes))
	}

	return nil
}

func (f *FindNoteGame) RunStep() error {
	// build answer
	answer := make(map[int][]*music.Note)

	for range f.NotesAmount {
		// pick a random note
		fretNumIdx := rand.Intn(len(f.Fretboard.Strings[0].FretNotes))

		note, err := f.Fretboard.GetNoteAt(0, fretNumIdx)
		if err != nil {
			return err
		}

		for range f.StringsAmount {
			// find the same note across a random string
			var allNotesInStr []*music.Note
			stringNum := rand.Intn(len(f.Fretboard.Strings) - 1)

			for _, currNote := range f.Fretboard.Strings[stringNum].FretNotes {
				if note.Name == currNote.Name && note.Symbol == currNote.Symbol {
					allNotesInStr = append(allNotesInStr, &currNote)
				}
			}

			answer[stringNum] = allNotesInStr
		}
	}

	// build question
	var sb strings.Builder
	sb.WriteString("Find note(s): [ ")
	for _, v := range answer {
		sb.WriteString(fmt.Sprintf("%s%s ", v[0].Name, v[0].Symbol))
	}
	sb.WriteString("]")
	sb.WriteString(" across string(s): [ ")
	for k, _ := range answer {
		sb.WriteString(fmt.Sprintf("%d ", k))
	}
	sb.WriteString("]")

	fmt.Println(sb.String())

	return nil
}

func (f *FindNoteGame) Summary() error {
	return nil
}

func (f *FindNoteGame) Quit() {
}
