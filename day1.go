package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Day1Part1 question:
// Your actual left and right lists contain many location IDs.
// What is the total distance between your lists?
func Day1Part1(puzzleInputName string) (string, error) {
	puzzleInput, err := GetPuzzleInput(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}

	puzzleInputLines := strings.Split(
		strings.Trim(puzzleInput, "\n"), "\n",
	)

	leftList := make([]int, len(puzzleInputLines))
	rightList := make([]int, len(puzzleInputLines))

	for i, line := range puzzleInputLines {
		var leftListValue int
		var rightListValue int

		_, err = fmt.Sscanf(line, "%d	%d", &leftListValue, &rightListValue)
		if err != nil {
			return stringErrorValue, err
		}

		leftList[i] = leftListValue
		rightList[i] = rightListValue
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := range puzzleInputLines {
		totalDistance += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}

	return strconv.Itoa(totalDistance), err
}
