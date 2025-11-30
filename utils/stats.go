package utils

import "fmt"

type Stats struct {
	totalQuestions int
	correctAnswers int
}

func NewStats() *Stats {
	return &Stats{
		totalQuestions: 0,
		correctAnswers: 0,
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

	fmt.Println("=======================")
	fmt.Println("[Game Stats			]")
	fmt.Printf("Num of questions: %d\n", s.totalQuestions)
	fmt.Printf("Correct Answers: %d\n", s.correctAnswers)
	fmt.Printf("Result: %d%%\n", correctPercentage)
	fmt.Println("========================")
}
