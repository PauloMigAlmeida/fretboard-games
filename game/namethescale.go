package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/PauloMigAlmeida/fretboard-games/utils"
	"io"
	"math/rand"
	"strings"
)

type NameTheScaleGame struct {
	StdIn  io.Reader
	StdOut io.Writer
	stats  *utils.Stats
	rng    *rand.Rand
}

func NewNameTheScaleGame(stdIn io.Reader, stdOut io.Writer, seed int64) *NameTheScaleGame {
	var rng *rand.Rand
	if seed != NoSeed {
		source := rand.NewSource(seed)
		rng = rand.New(source)
	} else {
		rng = rand.New(rand.NewSource(rand.Int63()))
	}

	return &NameTheScaleGame{
		StdIn:  stdIn,
		StdOut: stdOut,
		stats:  utils.NewStats(stdOut),
		rng:    rng,
	}
}

func (w *NameTheScaleGame) Configure() error {
	return nil
}

func (w *NameTheScaleGame) RunStep() error {
	allNotes := music.AllNotes()
	root := &allNotes[w.rng.Intn(len(allNotes))]

	allScales := music.AllScales()
	correctScaleIdx := w.rng.Intn(len(allScales))
	correctScale := allScales[correctScaleIdx]

	correctNotes, err := correctScale.NotesForRoot(root)
	if err != nil {
		return err
	}

	options, correctOptIdx := w.buildOptions(correctScaleIdx, allScales, root)

	w.Printf("Which scale contains those exact notes? [%s]\n", formatNotesSpaced(correctNotes))
	for i, opt := range options {
		w.Printf("  %d) %s\n", i+1, opt)
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
		w.Printf("Incorrect! ❌ - the correct answer was: %s\n", options[correctOptIdx])
		w.Printf("Formula for %s: %s\n", correctScale.Name, correctScale.Formula)
	}

	w.stats.RecordAnswer(correct)
	return nil
}

func (w *NameTheScaleGame) buildOptions(correctScaleIdx int, allScales []music.Scale, root *music.Note) ([]string, int) {
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

	correctLabel := fmt.Sprintf("%s%s %s", root.Name, root.Symbol, allScales[correctScaleIdx].Name)
	options := []string{correctLabel}

	for _, wi := range wrongIndices {
		label := fmt.Sprintf("%s%s %s", root.Name, root.Symbol, allScales[wi].Name)
		options = append(options, label)
	}

	w.rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	correctOptIdx := 0
	for i, opt := range options {
		if opt == correctLabel {
			correctOptIdx = i
			break
		}
	}

	return options, correctOptIdx
}

func (w *NameTheScaleGame) Summary() error {
	w.stats.PrintSummary()
	return nil
}

func (w *NameTheScaleGame) Quit() {
	_ = w.Summary()
}

func (w *NameTheScaleGame) Println(a ...any) {
	_, _ = fmt.Fprintln(w.StdOut, a...)
}

func (w *NameTheScaleGame) Print(a ...any) {
	_, _ = fmt.Fprint(w.StdOut, a...)
}

func (w *NameTheScaleGame) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(w.StdOut, format, a...)
}

func formatNotesSpaced(notes []*music.Note) string {
	parts := make([]string, len(notes))
	for i, n := range notes {
		parts[i] = n.PreferFlat()
	}
	return strings.Join(parts, " - ")
}
