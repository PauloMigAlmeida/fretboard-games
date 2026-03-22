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

	// seed 1234: string=3 (G), fret=11 → F#
	stdin.WriteString("F#\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Correct! ✅")
}

func TestWhatNoteIsItGame_RunStep_WhenEnharmonicAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	// F# and Gb are enharmonic equivalents
	stdin.WriteString("Gb\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Correct! ✅")
}

func TestWhatNoteIsItGame_RunStep_WhenWrongAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	stdin.WriteString("C\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Incorrect! ❌")
	assert.Contains(t, stdout.String(), "F#")
}

func TestWhatNoteIsItGame_RunStep_WhenInvalidAnswerIsGiven(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())
	var stdin, stdout bytes.Buffer

	game := NewWhatNoteIsItGame(fretboard, &stdin, &stdout, 1234)

	stdin.WriteString("XYZ\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Incorrect! ❌")
}
