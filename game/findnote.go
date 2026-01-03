package game

import (
	"fmt"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/PauloMigAlmeida/fretboard-games/utils"
	"io"
	"maps"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	NoSeed int64 = -1
)

type FindNoteGame struct {
	Fretboard *instrument.Fretboard
	// game variables
	NotesAmount   int
	StringsAmount int
	// OS stuff
	StdIn  io.Reader
	StdOut io.Writer
	// game misc
	stats *utils.Stats
	rng   *rand.Rand
}

func NewFindNoteGame(fretboard *instrument.Fretboard, stdIn io.Reader, stdOut io.Writer, seed int64) *FindNoteGame {
	var rng *rand.Rand
	if seed != NoSeed {
		source := rand.NewSource(seed)
		rng = rand.New(source)
	} else {
		rng = rand.New(rand.NewSource(rand.Int63()))
	}

	return &FindNoteGame{
		Fretboard:     fretboard,
		NotesAmount:   0,
		StringsAmount: 0,
		StdIn:         stdIn,
		StdOut:        stdOut,
		stats:         utils.NewStats(),
		rng:           rng,
	}
}

func (f *FindNoteGame) Configure() error {
	var notesAmount int
	f.Println("How many notes do you want to find?: ")
	_, err := fmt.Fscanf(f.StdIn, "%d\n", &notesAmount)
	if err != nil {
		return fmt.Errorf("error reading answer provider by user: %v", err)
	}
	f.NotesAmount = notesAmount

	var stringsAmount int
	f.Println("Across how many strings do you want to find notes?: ")
	_, err = fmt.Fscanf(f.StdIn, "%d\n", &stringsAmount)
	if err != nil {
		return fmt.Errorf("error reading answer provider by user: %v", err)
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
	answer, err := f.buildAnswer()
	if err != nil {
		return err
	}

	f.displayQuestionForAnswer(answer)

	userAnswer, err := f.readAnswerForQuestion(answer)
	if err != nil {
		return err
	}

	f.verifyAnswer(answer, userAnswer)

	//TODO interpret Contrl-C (probably no at the game level but at a higher abstraction level so it's reusable

	return nil
}

func (f *FindNoteGame) readAnswerForQuestion(answer map[int]map[int]*music.Note) (map[int]map[int]*music.Note, error) {
	ret := make(map[int]map[int]*music.Note, 0)

	stringNums := slices.Collect(maps.Keys(answer))
	sort.Ints(stringNums)

	for _, num := range stringNums {
		f.Printf("Enter answer for string [%d] (e.g., 3, 15): ", num)

		var userInput string
		_, err := fmt.Fscanf(f.StdIn, "%s\n", &userInput)
		if err != nil {
			return nil, fmt.Errorf("error reading answer provider by user: %v", err)
		}

		userAnswer, err := f.parseUserAnswer(userInput, num)
		if err != nil {
			return nil, err
		}

		maps.Copy(ret, userAnswer)
	}

	return ret, nil
}

func (f *FindNoteGame) buildAnswer() (map[int]map[int]*music.Note, error) {
	// sanity checks
	if len(f.Fretboard.Strings) < f.StringsAmount {
		return nil, fmt.Errorf("fretboard has less strings (%d) than the requested number of strings (%d)", len(f.Fretboard.Strings), f.StringsAmount)
	}

	if len(f.Fretboard.Strings[0].FretNotes) < f.NotesAmount {
		return nil, fmt.Errorf("fretboard has less frets (%d) than requested notes to find (%d)", len(f.Fretboard.Strings[0].FretNotes), f.NotesAmount)
	}

	randomStrings := make(map[int]bool)

	for len(randomStrings) < f.StringsAmount {
		stringNum := f.rng.Intn(len(f.Fretboard.Strings)-1) + 1

		if _, found := randomStrings[stringNum]; !found {
			randomStrings[stringNum] = true
		}
	}

	randomNotes := make(map[*music.Note]bool)

	for len(randomNotes) < f.NotesAmount {
		fretNumIdx := f.rng.Intn(len(f.Fretboard.Strings[0].FretNotes))

		// the string in itself isn't relevant here as we are trying to get the notes only
		note, err := f.Fretboard.GetNoteAt(1, fretNumIdx)
		if err != nil {
			return nil, err
		}

		if _, found := randomNotes[note]; !found {
			randomNotes[note] = true
		}
	}

	// map[stringNum]: { map[fretNum]: *note }
	answer := make(map[int]map[int]*music.Note)

	for stringNum := range randomStrings {
		combined := make(map[int]*music.Note)
		for note := range randomNotes {
			notesInStr := f.Fretboard.Strings[stringNum-1].FindNote(note)
			maps.Copy(combined, notesInStr)
		}
		answer[stringNum] = combined
	}

	return answer, nil
}

func (f *FindNoteGame) displayQuestionForAnswer(answer map[int]map[int]*music.Note) {
	seenNotes := map[*music.Note]bool{}

	f.Print("Find note(s) [ ")
	for _, str := range answer {
		for _, fret := range str {
			if !seenNotes[fret] {
				seenNotes[fret] = true
			}
		}
	}

	notes := slices.Collect(maps.Keys(seenNotes))
	sort.Slice(notes, func(i, j int) bool {
		return strings.Compare(
			fmt.Sprintf("%s%s", notes[i].Name, notes[i].Symbol),
			fmt.Sprintf("%s%s", notes[j].Name, notes[j].Symbol),
		) == -1
	})

	for _, note := range notes {
		f.Printf("%s%s ", note.Name, note.Symbol)
	}

	f.Print("] across string(s) [ ")

	stringNums := slices.Collect(maps.Keys(answer))
	slices.Sort(stringNums)

	for _, v := range stringNums {
		f.Printf("%d ", v)
	}
	f.Print("]\n")
}

func (f *FindNoteGame) parseUserAnswer(userAnswer string, stringNum int) (map[int]map[int]*music.Note, error) {
	answerMap := make(map[int]map[int]*music.Note)
	fretNumList := make([]int, 0)

	tokens := strings.Split(userAnswer, ",")
	for _, token := range tokens {
		fretNum, err := strconv.Atoi(strings.TrimSpace(token))

		if err != nil {
			return nil, fmt.Errorf("error parsing user-provider answer '%s': %v", token, err)
		}

		fretNumList = append(fretNumList, fretNum)
	}

	for _, fretNum := range fretNumList {
		note, err := f.Fretboard.GetNoteAt(stringNum, fretNum)

		if err != nil {
			return nil, fmt.Errorf("error note not found at fret number '%d': %v", fretNum, err)
		}

		if val, ok := answerMap[stringNum]; !ok {
			answerMap[stringNum] = map[int]*music.Note{
				fretNum: note,
			}
		} else {
			val[fretNum] = note
		}
	}

	return answerMap, nil
}

func (f *FindNoteGame) verifyAnswer(correctAnswer map[int]map[int]*music.Note, userAnswer map[int]map[int]*music.Note) {
	correct := true

	for caStrNum, caStrEl := range correctAnswer {
		uaStrEl, ok := userAnswer[caStrNum]

		if !ok {
			correct = false
			break
		}

		if len(caStrEl) != len(uaStrEl) {
			correct = false
			break
		}

		// compare
		for fretNum := range caStrEl {
			caNote, caFound := caStrEl[fretNum]
			uaNote, uaFound := uaStrEl[fretNum]

			if !caFound || !uaFound {
				correct = false
			} else {
				correct = correct && caNote.Equals(uaNote)
			}
		}
	}

	if correct {
		f.Println("Correct! ✅")
	} else {
		f.Println("Incorrect! ❌ - the correct answer was: [")
		f.displayFretboard(correctAnswer)
		f.Println("]")
	}

	f.stats.RecordAnswer(correct)
}

func (f *FindNoteGame) displayFretboard(correctAnswer map[int]map[int]*music.Note) {
	ignoredStrings := make([]int, 0)
	for strIdx := range f.Fretboard.Strings {
		if _, ok := correctAnswer[strIdx+1]; !ok {
			ignoredStrings = append(ignoredStrings, strIdx+1)
		}
	}

	notes := make([]*music.Note, 0)
	for _, fretNotes := range correctAnswer {
		for _, fretNote := range fretNotes {

			found := false
			for _, n := range notes {
				if n.Equals(fretNote) {
					found = true
				}
			}

			if !found {
				notes = append(notes, fretNote)
			}
		}
	}

	fretboardRepr, _ := f.Fretboard.DrawFretboard(notes, ignoredStrings)
	f.Println(fretboardRepr)
}

func (f *FindNoteGame) Summary() error {
	f.stats.PrintSummary()
	return nil
}

func (f *FindNoteGame) Quit() {
	_ = f.Summary()
}

func (f *FindNoteGame) Println(a ...any) {
	_, _ = fmt.Fprintln(f.StdOut, a...)
}

func (f *FindNoteGame) Print(a ...any) {
	_, _ = fmt.Fprint(f.StdOut, a...)
}

func (f *FindNoteGame) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(f.StdOut, format, a...)
}
