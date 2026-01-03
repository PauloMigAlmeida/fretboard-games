package utils

import (
	"fmt"
	"io"
)

type Stats struct {
	totalQuestions int
	correctAnswers int
	stdOut         io.Writer
}

func NewStats(stdOut io.Writer) *Stats {
	return &Stats{
		totalQuestions: 0,
		correctAnswers: 0,
		stdOut:         stdOut,
	}
}

func (s *Stats) RecordAnswer(correct bool) {
	s.totalQuestions++
	if correct {
		s.correctAnswers++
	}
}

func (s *Stats) PrintSummary() {
	var correctPercentage int
	if s.totalQuestions == 0 {
		correctPercentage = 0
	} else {
		correctPercentage = int((float64(s.correctAnswers) / float64(s.totalQuestions)) * 100)
	}

	s.Println("=======================")
	s.Println("[Game Stats]")
	s.Printf("Num of questions: %d\n", s.totalQuestions)
	s.Printf("Correct Answers: %d\n", s.correctAnswers)
	s.Printf("Result: %d%%\n", correctPercentage)
	s.Println("========================")
}

func (f *Stats) Println(a ...any) {
	_, _ = fmt.Fprintln(f.stdOut, a...)
}

func (f *Stats) Printf(format string, a ...any) {
	_, _ = fmt.Fprintf(f.stdOut, format, a...)
}
