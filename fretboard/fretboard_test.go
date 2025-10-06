package fretboard

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFretboard_GetNoteAt(t *testing.T) {
	// standard tuning test
	eNote, _ := FindNote(E, Natural)
	bNote, _ := FindNote(B, Natural)
	gNote, _ := FindNote(G, Natural)
	dNote, _ := FindNote(D, Natural)
	aNote, _ := FindNote(A, Natural)

	fretboard := NewFretboard(24, []*Note{
		eNote,
		bNote,
		gNote,
		dNote,
		aNote,
		eNote,
	})

	// test #1
	note, err := fretboard.GetNoteAt(1, 5)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(
		note,
		&Note{
			Name:            A,
			Symbol:          Natural,
			EnharmonicNames: []EnharmonicType{},
		}))

	// test #2
	note, err = fretboard.GetNoteAt(2, 5)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(
		note,
		&Note{
			Name:            E,
			Symbol:          Natural,
			EnharmonicNames: []EnharmonicType{{Name: F, Symbol: Flat}},
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
