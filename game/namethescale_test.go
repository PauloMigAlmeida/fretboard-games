package game

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameTheScaleGame_Configure(t *testing.T) {
	var stdin, stdout bytes.Buffer
	game := NewNameTheScaleGame(&stdin, &stdout, NoSeed)
	assert.Nil(t, game.Configure())
}

func TestNameTheScaleGame_RunStep_WhenCorrectAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewNameTheScaleGame(&stdin, &stdout, 1234)

	// seed 1234: root=D, scale=Melodic Minor, correct option is 2
	stdin.WriteString("2\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Correct! ✅")
}

func TestNameTheScaleGame_RunStep_WhenWrongAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewNameTheScaleGame(&stdin, &stdout, 1234)

	// seed 1234: root=D, scale=Melodic Minor, option 1 is wrong
	stdin.WriteString("1\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "Incorrect! ❌")
	assert.Contains(t, stdout.String(), "D Melodic Minor")
	assert.Contains(t, stdout.String(), "Formula for Melodic Minor: 1,2,b3,4,5,6,7")
}

func TestNameTheScaleGame_RunStep_WhenInvalidAnswerIsGiven(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewNameTheScaleGame(&stdin, &stdout, 1234)

	stdin.WriteString("9\n")
	err := game.RunStep()
	assert.NotNil(t, err)
}

func TestNameTheScaleGame_RunStep_QuestionContainsNotes(t *testing.T) {
	var stdin, stdout bytes.Buffer

	game := NewNameTheScaleGame(&stdin, &stdout, 1234)

	stdin.WriteString("2\n")
	err := game.RunStep()
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "D - E - F - G - A - B - Db")
}
