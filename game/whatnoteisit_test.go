package game

import (
	"bytes"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhatNoteIsItGame_Configure(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer
	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, NoSeed)
	assert.Nil(t, game.Configure())
}

func TestWhatNoteIsItGame_RunStep_WhenCorrectAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	// seed 1234: string=3, fret=11 → F#; after shuffling, correct option is 2
	stdin.WriteString("2\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Correct! ✅")
}

func TestWhatNoteIsItGame_RunStep_WhenWrongAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	// seed 1234: option 1 is E, which is wrong
	stdin.WriteString("1\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Incorrect! ❌")
	assert.Contains(t, stdout.String(), "F#")
}

func TestWhatNoteIsItGame_RunStep_WhenInvalidAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	stdin.WriteString("9\n")
	err := game.RunStep()
	assert.NotNil(t, err)
}
