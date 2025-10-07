package instrument

import (
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFretboard_GetNoteAt(t *testing.T) {
	fretboard := NewFretboard(24, StandardTuning())

	// test #1
	note, err := fretboard.GetNoteAt(1, 5)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(
		note,
		&music.Note{
			Name:            music.A,
			Symbol:          music.Natural,
			EnharmonicNames: []music.EnharmonicType{},
		}))

	// test #2
	note, err = fretboard.GetNoteAt(2, 5)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(
		note,
		&music.Note{
			Name:            music.E,
			Symbol:          music.Natural,
			EnharmonicNames: []music.EnharmonicType{{Name: music.F, Symbol: music.Flat}},
		}))

	// invalid string #1
	note, err = fretboard.GetNoteAt(7, 5)
	assert.Nil(t, note)
	assert.Error(t, err)

	// invalid string #2
	note, err = fretboard.GetNoteAt(0, 5)
	assert.Nil(t, note)
	assert.Error(t, err)

	// invalid fret #1
	note, err = fretboard.GetNoteAt(1, -1)
	assert.Nil(t, note)
	assert.Error(t, err)

	// invalid fret #2
	note, err = fretboard.GetNoteAt(1, 25)
	assert.Nil(t, note)
	assert.Error(t, err)
}
