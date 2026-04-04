package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/PauloMigAlmeida/fretboard-games/utils"
	"io"
	"math/rand"
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

	options := w.buildOptions(correctNote)
	correctIdx := w.indexOfCorrect(options, correctNote)

	w.Printf("What note is on fret %d of string %d?\n", fretNumber, stringNumber)
	for i, note := range options {
		w.Printf("  %d) %s%s\n", i+1, note.Name, note.Symbol)
	}
	w.Print("Your answer (1-3): ")

	var choice int
	_, err = fmt.Fscanf(w.StdIn, "%d\n", &choice)
	if err != nil || choice < 1 || choice > 3 {
		return fmt.Errorf("invalid input: expected a number between 1 and 3")
	}

	correct := choice-1 == correctIdx

	if correct {
		w.Println("Correct! ✅")
	} else {
		w.Printf("Incorrect! ❌ - the correct answer was: %s%s\n", correctNote.Name, correctNote.Symbol)
	}

	w.stats.RecordAnswer(correct)
	return nil
}

// buildOptions returns a shuffled slice of 3 notes: 1 correct + 2 distinct wrong ones.
func (w *WhatNoteIsItGame) buildOptions(correctNote *music.Note) []*music.Note {
	wrongMap := make(map[string]*music.Note, 2)
	for len(wrongMap) < 2 {
		candidate := music.RandomNote(w.rng)
		if candidate.Equals(correctNote) {
			continue
		}
		key := string(candidate.Name) + string(candidate.Symbol)
		if _, exists := wrongMap[key]; !exists {
			wrongMap[key] = candidate
		}
	}

	wrong := make([]*music.Note, 0, 2)
	for _, n := range wrongMap {
		wrong = append(wrong, n)
	}

	options := []*music.Note{correctNote, wrong[0], wrong[1]}
	w.rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	return options
}

func (w *WhatNoteIsItGame) indexOfCorrect(options []*music.Note, correctNote *music.Note) int {
	for i, n := range options {
		if n.Equals(correctNote) {
			return i
		}
	}
	return -1
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
