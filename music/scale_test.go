package music

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale_NotesForRoot_Major(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[MajorScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,D,E,F,G,A,B", FormatNotes(notes))
}

func TestScale_NotesForRoot_NaturalMinor(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[NaturalMinorScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,D,Eb,F,G,Ab,Bb", FormatNotes(notes))
}

func TestScale_NotesForRoot_MelodicMinor(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[MelodicMinorScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,D,Eb,F,G,A,B", FormatNotes(notes))
}

func TestScale_NotesForRoot_HarmonicMinor(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[HarmonicMinorScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,D,Eb,F,G,Ab,B", FormatNotes(notes))
}

func TestScale_NotesForRoot_Major_NonCRoot(t *testing.T) {
	// G Major: G A B C D E F#
	root, err := FindNote(G, Natural)
	assert.Nil(t, err)

	scale := AllScales()[MajorScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "G,A,B,C,D,E,Gb", FormatNotes(notes))
}

func TestNote_PreferFlat(t *testing.T) {
	cSharp, err := FindNote(C, Sharp)
	assert.Nil(t, err)
	assert.Equal(t, "Db", cSharp.PreferFlat())

	dSharp, err := FindNote(D, Sharp)
	assert.Nil(t, err)
	assert.Equal(t, "Eb", dSharp.PreferFlat())

	cNatural, err := FindNote(C, Natural)
	assert.Nil(t, err)
	assert.Equal(t, "C", cNatural.PreferFlat())
}

func TestAllScales_ReturnsAllFour(t *testing.T) {
	scales := AllScales()
	assert.Len(t, scales, 4)
	assert.Equal(t, "Major", scales[MajorScale].Name)
	assert.Equal(t, "Minor", scales[NaturalMinorScale].Name)
	assert.Equal(t, "Melodic Minor", scales[MelodicMinorScale].Name)
	assert.Equal(t, "Harmonic Minor", scales[HarmonicMinorScale].Name)
}
