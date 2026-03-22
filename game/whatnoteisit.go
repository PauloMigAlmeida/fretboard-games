package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/PauloMigAlmeida/fretboard-games/utils"
	"io"
	"math/rand"
	"strings"
)

type WhatNoteIsItGame struct {
	Fretboard *instrument.Fretboard
	StdIn     io.Reader
	StdOut    io.Writer
	stats     *utils.Stats
	rng       *rand.Rand
}

func NewWhatNoteIsItGame(fretboard *instrument.Fretboard, stdIn io.Reader, stdOut io.Writer, seed int64) *WhatNoteIsItGame {
	var rng *rand.Rand
	if seed != NoSeed {
		source := rand.NewSource(seed)
		rng = rand.New(source)
	} else {
		rng = rand.New(rand.NewSource(rand.Int63()))
	}

	return &WhatNoteIsItGame{
		Fretboard: fretboard,
		StdIn:     stdIn,
		StdOut:    stdOut,
		stats:     utils.NewStats(stdOut),
		rng:       rng,
	}
}

func (w *WhatNoteIsItGame) Configure() error {
	return nil
}

func (w *WhatNoteIsItGame) RunStep() error {
	stringNumber := w.rng.Intn(len(w.Fretboard.Strings)) + 1
	fretNumber := w.rng.Intn(len(w.Fretboard.Strings[0].FretNotes))

	correctNote, err := w.Fretboard.GetNoteAt(stringNumber, fretNumber)
	if err != nil {
		return err
	}

	w.Printf("What note is on fret %d of string %d?: ", fretNumber, stringNumber)

	var userInput string
	_, err = fmt.Fscanf(w.StdIn, "%s\n", &userInput)
	if err != nil {
		return fmt.Errorf("error reading answer provided by user: %v", err)
	}

	userInput = strings.TrimSpace(userInput)

	correct := w.checkAnswer(userInput, correctNote)

	if correct {
		w.Println("Correct! ✅")
	} else {
		w.Printf("Incorrect! ❌ - the correct answer was: %s%s\n", correctNote.Name, correctNote.Symbol)
	}

	w.stats.RecordAnswer(correct)
	return nil
}

func (w *WhatNoteIsItGame) checkAnswer(userInput string, correctNote *music.Note) bool {
	// parse user input into a Note so enharmonic equivalents are accepted
	var name music.NaturalNote
	var symbol music.Accidental

	switch len(userInput) {
	case 1:
		name = music.NaturalNote(strings.ToUpper(userInput))
		symbol = music.Natural
	case 2:
		name = music.NaturalNote(strings.ToUpper(string(userInput[0])))
		symbol = music.Accidental(string(userInput[1]))
	default:
		return false
	}

	userNote, err := music.FindNote(name, symbol)
	if err != nil {
		return false
	}

	return correctNote.Equals(userNote)
}

func (w *WhatNoteIsItGame) Summary() error {
	w.stats.PrintSummary()
	return nil
}

func (w *WhatNoteIsItGame) Quit() {
	_ = w.Summary()
}

func (w *WhatNoteIsItGame) Println(a ...any) {
	_, _ = fmt.Fprintln(w.StdOut, a...)
}

func (w *WhatNoteIsItGame) Print(a ...any) {
	_, _ = fmt.Fprint(w.StdOut, a...)
}

func (w *WhatNoteIsItGame) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(w.StdOut, format, a...)
}
