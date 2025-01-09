package main

import "testing"

const (
	part1CorrectAnswer = "11"
	part2CorrectAnswer = "31"
	puzzleInputName = "day1example"
)

func TestDay1Part1Example(t *testing.T) {
	answer, err := Day1Part1(puzzleInputName)
	if err != nil {
		t.Error(err)
	}

	if answer != part1CorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, part1CorrectAnswer)
	}
}

func TestDay1Part2Example(t *testing.T) {
	answer, err := Day1Part2(puzzleInputName)
	if err != nil {
		t.Error(err)
	}

	if answer != part2CorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, part2CorrectAnswer)
	}
}
