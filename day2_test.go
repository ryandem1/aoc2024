package main

import "testing"

const (
	day2PuzzleExampleInputName = "day2example"
	day2Part1ExampleCorrectAnswer = "2"
	day2Part2ExampleCorrectAnswer = "4"
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


func TestDay2Part2(t *testing.T) {
	answer, err := Day2Part2(day2PuzzleExampleInputName)
	if err != nil {
		t.Fatal(err)
	}

	if answer != day2Part2ExampleCorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, day2Part2ExampleCorrectAnswer)
	}
}
