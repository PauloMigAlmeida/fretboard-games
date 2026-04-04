package music

import (
	"fmt"
	"math/rand"
)

type NaturalNote string

const (
	C NaturalNote = "C"
	D NaturalNote = "D"
	E NaturalNote = "E"
	F NaturalNote = "F"
	G NaturalNote = "G"
	A NaturalNote = "A"
	B NaturalNote = "B"
)

type Accidental string

const (
	Natural Accidental = ""
	Sharp   Accidental = "#"
	Flat    Accidental = "b"
)

type EnharmonicType struct {
	Name   NaturalNote
	Symbol Accidental
}

type Note struct {
	Name            NaturalNote
	Symbol          Accidental
	EnharmonicNames []EnharmonicType
}

// Scratch pad:
//
//	b b  b b b
//
// w w ww w w w
var notes = []Note{
	{Name: C, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: B, Symbol: Sharp}}},
	{Name: C, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: D, Symbol: Flat}}},
	{Name: D, Symbol: Natural, EnharmonicNames: []EnharmonicType{}},
	{Name: D, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: E, Symbol: Flat}}},
	{Name: E, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: F, Symbol: Flat}}},
	{Name: F, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: E, Symbol: Sharp}}},
	{Name: F, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: G, Symbol: Flat}}},
	{Name: G, Symbol: Natural, EnharmonicNames: []EnharmonicType{}},
	{Name: G, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: A, Symbol: Flat}}},
	{Name: A, Symbol: Natural, EnharmonicNames: []EnharmonicType{}},
	{Name: A, Symbol: Sharp, EnharmonicNames: []EnharmonicType{{Name: B, Symbol: Flat}}},
	{Name: B, Symbol: Natural, EnharmonicNames: []EnharmonicType{{Name: C, Symbol: Flat}}},
}

// AllNotes returns all notes in the chromatic scale.
func AllNotes() []Note {
	return notes
}

// RandomNote returns a random note from the chromatic scale.
// If rng is nil, a new random source is used.
func RandomNote(rng *rand.Rand) *Note {
	if rng == nil {
		rng = rand.New(rand.NewSource(rand.Int63()))
	}
	return &notes[rng.Intn(len(notes))]
}

func FindNote(name NaturalNote, symbol Accidental) (*Note, error) {
	for _, note := range notes {
		if note.Name == name && note.Symbol == symbol {
			return &note, nil
		}

		for _, eNote := range note.EnharmonicNames {
			if eNote.Name == name && eNote.Symbol == symbol {
				return &note, nil
			}
		}
	}

	return nil, fmt.Errorf("note '%s%s' not found", name, symbol)
}

func (n *Note) NextHalfStepNote() (*Note, error) {
	for i, note := range notes {
		if n.Equals(&note) {
			if i == len(notes)-1 {
				return &notes[0], nil
			}
			return &notes[i+1], nil
		}
	}

	return nil, fmt.Errorf("note '%s' not found", n.Name)
}

func (n *Note) NextWholeStepNote() (*Note, error) {
	firstStep, err := n.NextHalfStepNote()
	if err != nil {
		return nil, err
	}
	return firstStep.NextHalfStepNote()
}

func (n *Note) PreviousHalfStepNote() (*Note, error) {
	for i, note := range notes {
		if n.Equals(&note) {
			if i == 0 {
				return &notes[len(notes)-1], nil
			}
			return &notes[i-1], nil
		}
	}
	return nil, fmt.Errorf("note '%s' not found", n.Name)
}

func (n *Note) PreviousWholeStepNote() (*Note, error) {
	firstStep, err := n.PreviousHalfStepNote()
	if err != nil {
		return nil, err
	}
	return firstStep.PreviousHalfStepNote()
}

func (n *Note) Equals(anotherNote *Note) bool {
	if n.Name == anotherNote.Name && n.Symbol == anotherNote.Symbol {
		return true
	}

	return false
}

// PreferFlat returns the note name using flat notation when the note is a sharp accidental
// (e.g., C# → Db, D# → Eb). Natural notes are returned as-is.
func (n *Note) PreferFlat() string {
	if n.Symbol == Sharp {
		for _, e := range n.EnharmonicNames {
			if e.Symbol == Flat {
				return string(e.Name) + string(e.Symbol)
			}
		}
	}
	return string(n.Name) + string(n.Symbol)
}
