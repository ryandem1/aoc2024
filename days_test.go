package main

import "testing"


func TestAllDays(t *testing.T) {
	testCases := []struct {
		name          string
		puzzleFunc    func(string) (string, error)
		puzzleInputName string
		correctAnswer string
	} {
		{
			name:          "Day1Part1Example",
			puzzleFunc:      Day1Part1,
			puzzleInputName: "day1example",
			correctAnswer: "11",
		},
		{
			name:          "Day1Part2Example",
			puzzleFunc:      Day1Part2,
			puzzleInputName: "day1example",
			correctAnswer: "31",
		},
		{
			name:          "Day2Part1Example",
			puzzleFunc:      Day2Part1,
			puzzleInputName: "day2example",
			correctAnswer: "2",
		},
		{
			name:          "Day2Part2Example",
			puzzleFunc:      Day2Part2,
			puzzleInputName: "day2example",
			correctAnswer: "4",
		},
		{
			name:          "Day3Part1Example",
			puzzleFunc:      Day3Part1,
			puzzleInputName: "day3example",
			correctAnswer: "161",
		},
		{
			name:          "Day3Part2Example",
			puzzleFunc:      Day3Part2,
			puzzleInputName: "day3example2",
			correctAnswer: "48",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			answer, err := tc.puzzleFunc(tc.puzzleInputName)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if answer != tc.correctAnswer {
				t.Errorf("answer: %s != expected: %s", answer, tc.correctAnswer)
			}
		})
	}
}

