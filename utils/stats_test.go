package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStats(t *testing.T) {
	stats := NewStats()
	assert.NotNil(t, stats)

	assert.Equal(t, 0, stats.totalQuestions)
	assert.Equal(t, 0, stats.correctAnswers)
}

func TestStats_RecordAnswer(t *testing.T) {
	stats := NewStats()
	assert.Equal(t, 0, stats.totalQuestions)
	assert.Equal(t, 0, stats.correctAnswers)

	stats.RecordAnswer(false)
	assert.Equal(t, 1, stats.totalQuestions)
	assert.Equal(t, 0, stats.correctAnswers)

	stats.RecordAnswer(true)
	assert.Equal(t, 2, stats.totalQuestions)
	assert.Equal(t, 1, stats.correctAnswers)
}

func TestStats_PrintSummary(t *testing.T) {
	stats := NewStats()
	assert.NotPanics(t, stats.PrintSummary)
}
