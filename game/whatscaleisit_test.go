package game

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhatScaleIsItGame_Configure(t *testing.T) {
	var stdin, stdout bytes.Buffer
	game := NewWhatScaleIsItGame(&stdin, &stdout, NoSeed)
	assert.Nil(t, game.Configure())
}

func TestWhatScaleIsItGame_RunStep_WhenCorrectAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewWhatScaleIsItGame(&stdin, &stdout, 1234)

	// seed 1234: root=D, scale=Harmonic Minor, correct option is 2
	stdin.WriteString("2\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Correct! ✅")
}

func TestWhatScaleIsItGame_RunStep_WhenWrongAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewWhatScaleIsItGame(&stdin, &stdout, 1234)

	// seed 1234: root=D, scale=Harmonic Minor, option 1 is wrong
	stdin.WriteString("1\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Incorrect! ❌")
	assert.Contains(t, stdout.String(), "D,E,F,G,A,Bb,Db")
	assert.Contains(t, stdout.String(), "Formula for Harmonic Minor: 1,2,b3,4,5,b6,7")
}

func TestWhatScaleIsItGame_RunStep_WhenInvalidAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewWhatScaleIsItGame(&stdin, &stdout, 1234)

	stdin.WriteString("9\n")
	err := game.RunStep()
	assert.NotNil(t, err)
}

func TestWhatScaleIsItGame_RunStep_QuestionContainsScaleName(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewWhatScaleIsItGame(&stdin, &stdout, 1234)

	stdin.WriteString("2\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "D Harmonic Minor scale")
}
