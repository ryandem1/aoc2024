package main

import "testing"

const ExampleCorrectAnswer = "11"

func TestDay1Part1Example(t *testing.T) {
	answer, err := Day1Part1("day1part1example")
	if err != nil {
		t.Error(err)
	}

	if answer != ExampleCorrectAnswer {
		t.Errorf("answer: %s != expected: %s", answer, ExampleCorrectAnswer)
	}
}
