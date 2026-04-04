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

func TestAllScales_ReturnsAllNine(t *testing.T) {
	scales := AllScales()
	assert.Len(t, scales, 9)
	assert.Equal(t, "Major", scales[MajorScale].Name)
	assert.Equal(t, "Minor", scales[NaturalMinorScale].Name)
	assert.Equal(t, "Melodic Minor", scales[MelodicMinorScale].Name)
	assert.Equal(t, "Harmonic Minor", scales[HarmonicMinorScale].Name)
	assert.Equal(t, "Major Pentatonic", scales[MajorPentatonicScale].Name)
	assert.Equal(t, "Minor Pentatonic", scales[MinorPentatonicScale].Name)
	assert.Equal(t, "Blues", scales[BluesScale].Name)
	assert.Equal(t, "Spanish", scales[SpanishScale].Name)
	assert.Equal(t, "Persian", scales[PersianScale].Name)
}

func TestScale_NotesForRoot_MajorPentatonic(t *testing.T) {
	root, err := FindNote(A, Natural)
	assert.Nil(t, err)

	scale := AllScales()[MajorPentatonicScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "A,B,Db,E,Gb", FormatNotes(notes))
}

func TestScale_NotesForRoot_MinorPentatonic(t *testing.T) {
	root, err := FindNote(A, Natural)
	assert.Nil(t, err)

	scale := AllScales()[MinorPentatonicScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "A,C,D,E,G", FormatNotes(notes))
}

func TestScale_NotesForRoot_Blues(t *testing.T) {
	root, err := FindNote(A, Natural)
	assert.Nil(t, err)

	scale := AllScales()[BluesScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "A,C,D,Eb,E,G", FormatNotes(notes))
}

func TestScale_NotesForRoot_Spanish(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[SpanishScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,Db,E,F,G,Ab,Bb", FormatNotes(notes))
}

func TestScale_NotesForRoot_Persian(t *testing.T) {
	root, err := FindNote(C, Natural)
	assert.Nil(t, err)

	scale := AllScales()[PersianScale]
	notes, err := scale.NotesForRoot(root)
	assert.Nil(t, err)
	assert.Equal(t, "C,Db,E,F,Gb,Ab,B", FormatNotes(notes))
}
