package game

import (
	"bytes"
	"github.com/PauloMigAlmeida/fretboard-games/instrument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNoteGame_Configure(t *testing.T) {
	fretboard := instrument.NewFretboard(24, instrument.StandardTuning())

	var stdin bytes.Buffer
	stdin.WriteString("1\n2\n")

	game := NewFindNoteGame(fretboard, &stdin)

	err := game.Configure()
	assert.Nil(t, err)
	assert.Equal(t, 1, game.NotesAmount)
	assert.Equal(t, 2, game.StringsAmount)
}
