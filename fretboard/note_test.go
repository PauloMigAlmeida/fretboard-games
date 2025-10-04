package fretboard

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNote_TestFindNote(t *testing.T) {
	// simple note
	note, err := FindNote(D, Natural)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(note, &Note{
		D,
		Natural,
		[]EnharmonicType{},
	}))

	// note with enharmonic names
	note, err = FindNote(C, Sharp)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(note, &Note{
		C,
		Sharp,
		[]EnharmonicType{{Name: D, Symbol: Flat}},
	}))

	// doesn't directly exist as a note, but it's found as an enharmonic note name
	note, err = FindNote(E, Sharp)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(note, &Note{
		F,
		Natural,
		[]EnharmonicType{{Name: E, Symbol: Sharp}},
	}))
}

func TestNote_NextHalfStepNote(t *testing.T) {
	// simple half step
	note, _ := FindNote(D, Natural)
	nextHalfStep, err := note.NextHalfStepNote()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(nextHalfStep, &Note{
		D,
		Sharp,
		[]EnharmonicType{{Name: E, Symbol: Flat}},
	}))

	// half step the triggers circular queue logic
	note, _ = FindNote(B, Natural)
	nextHalfStep, err = note.NextHalfStepNote()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(nextHalfStep, &Note{
		C,
		Natural,
		[]EnharmonicType{{Name: B, Symbol: Sharp}},
	}))
}

func TestNote_PreviousHalfStepNote(t *testing.T) {
	// simple half step
	note, _ := FindNote(D, Natural)
	nextHalfStep, err := note.PreviousHalfStepNote()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(nextHalfStep, &Note{
		C,
		Sharp,
		[]EnharmonicType{{Name: D, Symbol: Flat}},
	}))

	// half step the triggers circular queue logic
	note, _ = FindNote(C, Natural)
	nextHalfStep, err = note.PreviousHalfStepNote()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(nextHalfStep, &Note{
		B,
		Natural,
		[]EnharmonicType{{Name: C, Symbol: Flat}},
	}))
}
