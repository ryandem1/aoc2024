package main

import "testing"

const (
	day2PuzzleExampleInputName = "day2example"
	day2Part1ExampleCorrectAnswer = "2"
)

func TestDay2Part1(t *testing.T) {
	answer, err := Day2Part1(day2PuzzleExampleInputName)
	if err != nil {
		t.Fatal(err)
	}

	if answer != day2Part1ExampleCorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, day2Part1ExampleCorrectAnswer)
	}
}