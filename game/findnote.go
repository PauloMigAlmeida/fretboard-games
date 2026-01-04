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
		stats:         utils.NewStats(stdOut),
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
	correctAnswer, err := f.buildAnswer()
	if err != nil {
		return err
	}

	f.displayQuestionForAnswer(correctAnswer)

	userSubmittedAnswer, err := f.readAnswerForQuestion(correctAnswer)
	if err != nil {
		return err
	}

	f.verifyAnswer(correctAnswer, userSubmittedAnswer)

	return nil
}

func (f *FindNoteGame) readAnswerForQuestion(correctAnswer map[int]map[int]*music.Note) (map[int]map[int]*music.Note, error) {
	userSubmittedAnswer := make(map[int]map[int]*music.Note, 0)

	stringNumbers := slices.Collect(maps.Keys(correctAnswer))
	sort.Ints(stringNumbers)

	for _, stringNumber := range stringNumbers {
		f.Printf("Enter answer for string [%d] (e.g., 3, 15): ", stringNumber)

		var userInput string
		_, err := fmt.Fscanf(f.StdIn, "%s\n", &userInput)
		if err != nil {
			return nil, fmt.Errorf("error reading answer provider by user: %v", err)
		}

		parsedAnswer, err := f.parseUserAnswer(userInput, stringNumber)
		if err != nil {
			return nil, err
		}

		maps.Copy(userSubmittedAnswer, parsedAnswer)
	}

	return userSubmittedAnswer, nil
}

func (f *FindNoteGame) buildAnswer() (map[int]map[int]*music.Note, error) {
	// sanity checks
	if len(f.Fretboard.Strings) < f.StringsAmount {
		return nil, fmt.Errorf("fretboard has less strings (%d) than the requested number of strings (%d)", len(f.Fretboard.Strings), f.StringsAmount)
	}

	if len(f.Fretboard.Strings[0].FretNotes) < f.NotesAmount {
		return nil, fmt.Errorf("fretboard has less frets (%d) than requested notes to find (%d)", len(f.Fretboard.Strings[0].FretNotes), f.NotesAmount)
	}

	selectedStrings := make(map[int]bool)

	for len(selectedStrings) < f.StringsAmount {
		stringNumber := f.rng.Intn(len(f.Fretboard.Strings)-1) + 1

		if _, alreadySelected := selectedStrings[stringNumber]; !alreadySelected {
			selectedStrings[stringNumber] = true
		}
	}

	targetNotes := make(map[*music.Note]bool)

	for len(targetNotes) < f.NotesAmount {
		fretIndex := f.rng.Intn(len(f.Fretboard.Strings[0].FretNotes))

		// the string in itself isn't relevant here as we are trying to get the notes only
		note, err := f.Fretboard.GetNoteAt(1, fretIndex)
		if err != nil {
			return nil, err
		}

		if _, alreadySelected := targetNotes[note]; !alreadySelected {
			targetNotes[note] = true
		}
	}

	// map[stringNumber]: { map[fretNumber]: *note }
	gameAnswer := make(map[int]map[int]*music.Note)

	for stringNumber := range selectedStrings {
		fretPositionsForString := make(map[int]*music.Note)
		for note := range targetNotes {
			fretPositions := f.Fretboard.Strings[stringNumber-1].FindNote(note)
			maps.Copy(fretPositionsForString, fretPositions)
		}
		gameAnswer[stringNumber] = fretPositionsForString
	}

	return gameAnswer, nil
}

func (f *FindNoteGame) displayQuestionForAnswer(correctAnswer map[int]map[int]*music.Note) {
	uniqueNotes := map[*music.Note]bool{}

	f.Print("Find note(s) [ ")
	for _, stringFrets := range correctAnswer {
		for _, noteAtFret := range stringFrets {
			if !uniqueNotes[noteAtFret] {
				uniqueNotes[noteAtFret] = true
			}
		}
	}

	sortedNotes := slices.Collect(maps.Keys(uniqueNotes))
	sort.Slice(sortedNotes, func(i, j int) bool {
		return strings.Compare(
			fmt.Sprintf("%s%s", sortedNotes[i].Name, sortedNotes[i].Symbol),
			fmt.Sprintf("%s%s", sortedNotes[j].Name, sortedNotes[j].Symbol),
		) == -1
	})

	for _, note := range sortedNotes {
		f.Printf("%s%s ", note.Name, note.Symbol)
	}

	f.Print("] across string(s) [ ")

	stringNumbers := slices.Collect(maps.Keys(correctAnswer))
	slices.Sort(stringNumbers)

	for _, stringNumber := range stringNumbers {
		f.Printf("%d ", stringNumber)
	}
	f.Print("]\n")
}

func (f *FindNoteGame) parseUserAnswer(userInputString string, stringNumber int) (map[int]map[int]*music.Note, error) {
	parsedAnswerMap := make(map[int]map[int]*music.Note)
	fretNumbersList := make([]int, 0)

	inputTokens := strings.Split(userInputString, ",")
	for _, token := range inputTokens {
		fretNumber, err := strconv.Atoi(strings.TrimSpace(token))

		if err != nil {
			return nil, fmt.Errorf("error parsing user-provider answer '%s': %v", token, err)
		}

		fretNumbersList = append(fretNumbersList, fretNumber)
	}

	for _, fretNumber := range fretNumbersList {
		note, err := f.Fretboard.GetNoteAt(stringNumber, fretNumber)

		if err != nil {
			return nil, fmt.Errorf("error note not found at fret number '%d': %v", fretNumber, err)
		}

		if existingFrets, stringExists := parsedAnswerMap[stringNumber]; !stringExists {
			parsedAnswerMap[stringNumber] = map[int]*music.Note{
				fretNumber: note,
			}
		} else {
			existingFrets[fretNumber] = note
		}
	}

	return parsedAnswerMap, nil
}

func (f *FindNoteGame) verifyAnswer(correctAnswer map[int]map[int]*music.Note, userAnswer map[int]map[int]*music.Note) {
	isAnswerCorrect := true

	for stringNumber, correctStringFrets := range correctAnswer {
		userStringFrets, stringExistsInUserAnswer := userAnswer[stringNumber]

		if !stringExistsInUserAnswer {
			isAnswerCorrect = false
			break
		}

		if len(correctStringFrets) != len(userStringFrets) {
			isAnswerCorrect = false
			break
		}

		// compare each fret position
		for fretNumber := range correctStringFrets {
			correctNote, correctNoteFound := correctStringFrets[fretNumber]
			userNote, userNoteFound := userStringFrets[fretNumber]

			if !correctNoteFound || !userNoteFound {
				isAnswerCorrect = false
			} else {
				isAnswerCorrect = isAnswerCorrect && correctNote.Equals(userNote)
			}
		}
	}

	if isAnswerCorrect {
		f.Println("Correct! ✅")
	} else {
		f.Println("Incorrect! ❌ - the correct answer was: [")
		f.displayFretboard(correctAnswer)
		f.Println("]")
	}

	f.stats.RecordAnswer(isAnswerCorrect)
}

func (f *FindNoteGame) displayFretboard(correctAnswer map[int]map[int]*music.Note) {
	stringsToIgnore := make([]int, 0)
	for stringIndex := range f.Fretboard.Strings {
		if _, isStringInAnswer := correctAnswer[stringIndex+1]; !isStringInAnswer {
			stringsToIgnore = append(stringsToIgnore, stringIndex+1)
		}
	}

	uniqueNotesToHighlight := make([]*music.Note, 0)
	for _, fretPositions := range correctAnswer {
		for _, noteAtPosition := range fretPositions {

			alreadyAdded := false
			for _, existingNote := range uniqueNotesToHighlight {
				if existingNote.Equals(noteAtPosition) {
					alreadyAdded = true
				}
			}

			if !alreadyAdded {
				uniqueNotesToHighlight = append(uniqueNotesToHighlight, noteAtPosition)
			}
		}
	}

	fretboardVisualization, _ := f.Fretboard.DrawFretboard(uniqueNotesToHighlight, stringsToIgnore)
	f.Println(fretboardVisualization)
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
