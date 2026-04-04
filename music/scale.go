package music

import (
	"fmt"
	"strings"
)

type ScaleType int

const (
	MajorScale ScaleType = iota
	NaturalMinorScale
	MelodicMinorScale
	HarmonicMinorScale
)

type Scale struct {
	Type      ScaleType
	Name      string
	Intervals []int  // semitones from root
	Formula   string // e.g., "1,2,b3,4,5,b6,b7"
}

var scales = []Scale{
	{
		Type:      MajorScale,
		Name:      "Major",
		Intervals: []int{0, 2, 4, 5, 7, 9, 11},
		Formula:   "1,2,3,4,5,6,7",
	},
	{
		Type:      NaturalMinorScale,
		Name:      "Minor",
		Intervals: []int{0, 2, 3, 5, 7, 8, 10},
		Formula:   "1,2,b3,4,5,b6,b7",
	},
	{
		Type:      MelodicMinorScale,
		Name:      "Melodic Minor",
		Intervals: []int{0, 2, 3, 5, 7, 9, 11},
		Formula:   "1,2,b3,4,5,6,7",
	},
	{
		Type:      HarmonicMinorScale,
		Name:      "Harmonic Minor",
		Intervals: []int{0, 2, 3, 5, 7, 8, 11},
		Formula:   "1,2,b3,4,5,b6,7",
	},
}

// AllScales returns all available scale definitions.
func AllScales() []Scale {
	return scales
}

// NotesForRoot computes the notes of this scale starting from the given root note.
func (s *Scale) NotesForRoot(root *Note) ([]*Note, error) {
	result := make([]*Note, len(s.Intervals))
	for i, semitones := range s.Intervals {
		n := root
		for step := 0; step < semitones; step++ {
			next, err := n.NextHalfStepNote()
			if err != nil {
				return nil, fmt.Errorf("error computing scale note: %w", err)
			}
			n = next
		}
		result[i] = n
	}
	return result, nil
}

// FormatNotes formats a slice of notes as a comma-separated string, preferring flat notation.
func FormatNotes(notes []*Note) string {
	parts := make([]string, len(notes))
	for i, n := range notes {
		parts[i] = n.PreferFlat()
	}
	return strings.Join(parts, ",")
}
