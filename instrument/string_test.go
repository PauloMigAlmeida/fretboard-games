package instrument

import (
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewString(t *testing.T) {
	eNote, _ := music.FindNote(music.E, music.Natural)
	str := NewString(eNote, 5)
	assert.Len(t, str.FretNotes, 5)

	expected := [5]music.Note{
		{Name: music.E, Symbol: music.Natural, EnharmonicNames: []music.EnharmonicType{{Name: music.F, Symbol: music.Flat}}},
		{Name: music.F, Symbol: music.Natural, EnharmonicNames: []music.EnharmonicType{{Name: music.E, Symbol: music.Sharp}}},
		{Name: music.F, Symbol: music.Sharp, EnharmonicNames: []music.EnharmonicType{{Name: music.G, Symbol: music.Flat}}},
		{Name: music.G, Symbol: music.Natural, EnharmonicNames: []music.EnharmonicType{}},
		{Name: music.G, Symbol: music.Sharp, EnharmonicNames: []music.EnharmonicType{{Name: music.A, Symbol: music.Flat}}},
	}

	for i := range 5 {
		assert.True(t, reflect.DeepEqual(str.FretNotes[i], expected[i]))
	}
}

func TestString_FindNote(t *testing.T) {
	eNote, _ := music.FindNote(music.E, music.Natural)
	str := NewString(eNote, 24)

	assert.Equal(t, str.FindNote(eNote), map[int]*music.Note{
		0:  eNote,
		12: eNote,
	})
}
