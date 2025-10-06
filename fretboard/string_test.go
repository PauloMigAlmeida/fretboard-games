package fretboard

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewString(t *testing.T) {
	eNote, _ := FindNote(E, Natural)
	str := NewString(eNote, 5)
	assert.Len(t, str.FretNotes, 6) // open note is fret "0"

	expected := [6]Note{
		{Name: E, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: F, Symbol: Flat}}},
		{Name: F, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: E, Symbol: Sharp}}},
		{Name: F, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: G, Symbol: Flat}}},
		{Name: G, Symbol: Natural, EnharmonicNames: []EnharmonicType{}},
		{Name: G, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: A, Symbol: Flat}}},
		{Name: A, Symbol: Natural, EnharmonicNames: []EnharmonicType{}},
	}

	for i := range 6 {
		assert.True(t, reflect.DeepEqual(str.FretNotes[i], expected[i]))
	}
}
