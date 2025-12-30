package instrument

import (
	"github.com/PauloMigAlmeida/fretboard-games/music"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
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

func TestFretboard_DrawFretboard(t *testing.T) {
	invalidFretboard := NewFretboard(0, StandardTuning())

	_, err := invalidFretboard.DrawFretboard([]*music.Note{}, []int{})
	assert.NotNil(t, err)

	validFretboard := NewFretboard(12, StandardTuning())

	noteA, _ := music.FindNote(music.A, music.Natural)
	noteC, _ := music.FindNote(music.C, music.Natural)

	// Without ignoring strings
	ret, err := validFretboard.DrawFretboard([]*music.Note{
		noteA,
	}, []int{})
	assert.Nil(t, err)

	assert.Equal(t, strings.TrimSpace(`
| 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | 10 | 11 |
| -  | -  | -  | -  | -  | X  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | X  | -  |
| -  | -  | X  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | X  | -  | -  | -  | -  |
| X  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | X  | -  | -  | -  | -  | -  | -  |
`), ret)

	ret, err = validFretboard.DrawFretboard([]*music.Note{
		noteC,
	}, []int{})
	assert.Nil(t, err)

	assert.Equal(t, strings.TrimSpace(`
| 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | 10 | 11 |
| -  | -  | -  | -  | -  | -  | -  | -  | X  | -  | -  | -  |
| -  | X  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | X  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | X  | -  |
| -  | -  | -  | X  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | X  | -  | -  | -  |
`), ret)

	// ignoring strings
	ret, err = validFretboard.DrawFretboard([]*music.Note{
		noteA,
	}, []int{
		6, 5, 4, 3, 2,
	})
	assert.Nil(t, err)

	assert.Equal(t, strings.TrimSpace(`
| 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | 10 | 11 |
| -  | -  | -  | -  | -  | X  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | X  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
| -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  | -  |
`), ret)
}
