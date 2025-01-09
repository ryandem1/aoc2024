package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)


// parseLeftAndRightList will parse the left and right list of
// location IDs from the puzzle input.
func parseLeftAndRightList(puzzleInputName string) ([]int, []int, error) {
	puzzleInputLines, err := GetPuzzleInputByLines(puzzleInputName)
	if err != nil {
		return nil, nil, err
	}

	leftList := make([]int, len(puzzleInputLines))
	rightList := make([]int, len(puzzleInputLines))

	for i, line := range puzzleInputLines {
		var leftListValue int
		var rightListValue int

		_, err = fmt.Sscanf(line, "%d	%d", &leftListValue, &rightListValue)
		if err != nil {
			return nil, nil, err
		}

		leftList[i] = leftListValue
		rightList[i] = rightListValue
	}

	return leftList, rightList, nil
}

func Day1Part1(puzzleInputName string) (string, error) {
	leftList, rightList, err := parseLeftAndRightList(puzzleInputName)

	if err != nil {
		return stringErrorValue, err
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := range len(leftList) {
		totalDistance += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}

	return strconv.Itoa(totalDistance), err
}

func Day1Part2(puzzleInputName string) (string, error) {
	leftList, rightList, err := parseLeftAndRightList(puzzleInputName)

	if err != nil {
		return stringErrorValue, err
	}

	rightListLocationIDFreq := make(map[int]int)
	for _, locationID := range rightList {
		if _, ok := rightListLocationIDFreq[locationID]; ok {
			rightListLocationIDFreq[locationID]++
		} else {
			rightListLocationIDFreq[locationID] = 1
		}
	}

	similarityScore := 0
	for _, locationID := range leftList {
		if freq, ok := rightListLocationIDFreq[locationID]; ok {
			similarityScore += locationID * freq
		}
	}

	return strconv.Itoa(similarityScore), nil
}
