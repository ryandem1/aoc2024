package main

import "testing"

const (
	day1Part1CorrectAnswer = "11"
	day1Part2CorrectAnswer = "31"
	day1PuzzleInputName = "day1example"
)

func TestDay1Part1Example(t *testing.T) {
	answer, err := Day1Part1(day1PuzzleInputName)
	if err != nil {
		t.Error(err)
	}

	if answer != day1Part1CorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, day1Part1CorrectAnswer)
	}
}

func TestDay1Part2Example(t *testing.T) {
	answer, err := Day1Part2(day1PuzzleInputName)
	if err != nil {
		t.Error(err)
	}

	if answer != day1Part2CorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, day1Part2CorrectAnswer)
	}
}
