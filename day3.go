package main

import (
	"fmt"
	"github.com/ryandem1/aoc2024/common"
	"strconv"
)

func Day3Part1(puzzleInputName string) (string, error) {
	puzzleInput, err := common.GetPuzzleInput(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}

	var multiplicand int
	var multiplier int

	total := 0
	for i := range puzzleInput {

		// so the sscan can find 'mul(%d,%d' if it is at the beginning of the string.
		// the problem is that it doesn't guarentee the sequence will end in ')', it is going
		// to ignore everything after it finds the first 2 digits.

		// so, to ensure it is a valid mul command, you have to check for the closing ')'
		numMatched, _ := fmt.Sscanf(puzzleInput[i:], "mul(%d,%d)", &multiplicand, &multiplier)
		if numMatched == 2 {
			numDigits := common.CountDigits(multiplicand) + common.CountDigits(multiplier)
			if puzzleInput[i + numDigits + 5] == ')' {
				total += multiplicand * multiplier
			}
		}
	}
	return strconv.Itoa(total), nil
}


func Day3Part2(puzzleInputName string) (string, error) {
	return puzzleInputName, nil
}
