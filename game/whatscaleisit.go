package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/PauloMigAlmeida/fretboard-games/utils"
	"io"
	"math/rand"
)

type WhatScaleIsItGame struct {
	StdIn  io.Reader
	StdOut io.Writer
	stats  *utils.Stats
	rng    *rand.Rand
}

func NewWhatScaleIsItGame(stdIn io.Reader, stdOut io.Writer, seed int64) *WhatScaleIsItGame {
	var rng *rand.Rand
	if seed != NoSeed {
		source := rand.NewSource(seed)
		rng = rand.New(source)
	} else {
		rng = rand.New(rand.NewSource(rand.Int63()))
	}

	return &WhatScaleIsItGame{
		StdIn:  stdIn,
		StdOut: stdOut,
		stats:  utils.NewStats(stdOut),
		rng:    rng,
	}
}

func (w *WhatScaleIsItGame) Configure() error {
	return nil
}

func (w *WhatScaleIsItGame) RunStep() error {
	allNotes := music.AllNotes()
	root := &allNotes[w.rng.Intn(len(allNotes))]

	allScales := music.AllScales()
	correctScaleIdx := w.rng.Intn(len(allScales))
	correctScale := allScales[correctScaleIdx]

	correctNotes, err := correctScale.NotesForRoot(root)
	if err != nil {
		return err
	}

	options, correctOptIdx := w.buildOptions(root, correctScaleIdx, allScales, correctNotes)

	w.Printf("What notes are part of the %s%s %s scale?\n", root.Name, root.Symbol, correctScale.Name)
	for i, opt := range options {
		w.Printf("  %d) %s\n", i+1, opt.formatted)
	}
	w.Print("Your answer (1-3): ")

	var choice int
	_, err = fmt.Fscanf(w.StdIn, "%d\n", &choice)
	if err != nil || choice < 1 || choice > 3 {
		return fmt.Errorf("invalid input: expected a number between 1 and 3")
	}

	correct := choice-1 == correctOptIdx

	if correct {
		w.Println("Correct! ✅")
	} else {
		w.Printf("Incorrect! ❌ - the correct answer was: %s\n", options[correctOptIdx].formatted)
		w.Printf("Formula for %s: %s\n", correctScale.Name, correctScale.Formula)
	}

	w.stats.RecordAnswer(correct)
	return nil
}

type scaleOption struct {
	formatted string
}

func (w *WhatScaleIsItGame) buildOptions(root *music.Note, correctScaleIdx int, allScales []music.Scale, correctNotes []*music.Note) ([]scaleOption, int) {
	// Pick 2 distinct wrong scale indices (different from the correct one)
	wrongIndices := make([]int, 0, 2)
	for len(wrongIndices) < 2 {
		idx := w.rng.Intn(len(allScales))
		if idx == correctScaleIdx {
			continue
		}
		alreadyPicked := false
		for _, wi := range wrongIndices {
			if wi == idx {
				alreadyPicked = true
				break
			}
		}
		if !alreadyPicked {
			wrongIndices = append(wrongIndices, idx)
		}
	}

	options := []scaleOption{
		{formatted: music.FormatNotes(correctNotes)},
	}

	for _, wi := range wrongIndices {
		wrongNotes, _ := allScales[wi].NotesForRoot(root)
		options = append(options, scaleOption{formatted: music.FormatNotes(wrongNotes)})
	}

	// Shuffle options, track where the correct one ends up
	w.rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	correctFormatted := music.FormatNotes(correctNotes)
	correctOptIdx := 0
	for i, opt := range options {
		if opt.formatted == correctFormatted {
			correctOptIdx = i
			break
		}
	}

	return options, correctOptIdx
}

func (w *WhatScaleIsItGame) Summary() error {
	w.stats.PrintSummary()
	return nil
}

func (w *WhatScaleIsItGame) Quit() {
	_ = w.Summary()
}

func (w *WhatScaleIsItGame) Println(a ...any) {
	_, _ = fmt.Fprintln(w.StdOut, a...)
}

func (w *WhatScaleIsItGame) Print(a ...any) {
	_, _ = fmt.Fprint(w.StdOut, a...)
}

func (w *WhatScaleIsItGame) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(w.StdOut, format, a...)
}
