package game

import (
	"bytes"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindNoteGame_Configure(t *testing.T) {
	// happy path
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())

	var stdin bytes.Buffer
	stdin.WriteString("1\n2\n")
	var stdout bytes.Buffer

	game := NewFindNoteGame(fretboard, &stdin, &stdout, NoSeed)

	err := game.Configure()
	assert.Nil(t, err)
	assert.Equal(t, 1, game.NotesAmount)
	assert.Equal(t, 2, game.StringsAmount)

	// invalid input
	stdin.Reset()
	stdout.Reset()
	stdin.WriteString("a\n2\n")
	game = NewFindNoteGame(fretboard, &stdin, &stdout, NoSeed)

	err = game.Configure()
	assert.NotNil(t, err)
}

func TestFindNoteGame_RunStep_WhenCorrectAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin bytes.Buffer
	var stdout bytes.Buffer

	game := NewFindNoteGame(fretboard, &stdin, &stdout, 1234)
	game.NotesAmount = 1
	game.StringsAmount = 1

	stdin.WriteString("2,14\n")
	err := game.RunStep()
	assert.Nil(t, err)

	buf, _ := game.StdOut.(*bytes.Buffer)
	bufStr := buf.String()
	assert.Contains(t, bufStr, "Find note(s) [ F# ] across string(s) [ 1 ]")
	assert.Contains(t, bufStr, "Correct! ✅")
}

func TestFindNoteGame_RunStep_WhenWrongAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin bytes.Buffer
	var stdout bytes.Buffer

	game := NewFindNoteGame(fretboard, &stdin, &stdout, 1234)
	game.NotesAmount = 1
	game.StringsAmount = 1

	stdin.WriteString("2,13\n")
	err := game.RunStep()
	assert.Nil(t, err)

	buf, _ := game.StdOut.(*bytes.Buffer)
	bufStr := buf.String()
	assert.Contains(t, bufStr, "Find note(s) [ F# ] across string(s) [ 1 ]")
	assert.Contains(t, bufStr, "Incorrect! ❌")
	assert.Contains(t, bufStr, strings.TrimSpace(`
| 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
| -  | -  | X  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | X  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
`))
}

func TestFindNoteGame_RunStep_WithMultipleNotes_WhenCorrectAnswerIsGiven(t *testing.T) {
	//TODO this test (and the logic behind it, needs to be tweaked)
	t.Skip("this test (and the logic behind it, needs to be tweaked)")

	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin bytes.Buffer
	var stdout bytes.Buffer

	game := NewFindNoteGame(fretboard, &stdin, &stdout, 1234)
	game.NotesAmount = 3
	game.StringsAmount = 3

	stdin.WriteString("9,21,1,13,8,20,")
	stdin.WriteString("0,12,4,16,11,23,")
	stdin.WriteString("7,19,11,23,6,18")
	err := game.RunStep()
	assert.Nil(t, err)

	buf, _ := game.StdOut.(*bytes.Buffer)
	bufStr := buf.String()
	assert.Contains(t, bufStr, "Find note(s) [ D# E G# ] across string(s) [ 1 3 5 ]")
	assert.Contains(t, bufStr, "Correct! ✅")
}
