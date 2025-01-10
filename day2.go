package main

import (
	"github.com/ryandem1/aoc2024/common"
	"math"
	"strconv"
	"strings"
)

func parseReportsFromPuzzleInput(puzzleInputName string) ([][]int, error) {
	puzzleInputLines, err := common.GetPuzzleInputByLines(puzzleInputName)
	if err != nil {
		return nil, err
	}

	reports := make([][]int, len(puzzleInputLines))
	for i, line := range puzzleInputLines {
		stringLevels := strings.Split(line, " ")
		report := make([]int, len(stringLevels))

		for j, stringLevel := range stringLevels {
			level, err := strconv.Atoi(string(stringLevel))
			if err != nil {
				return nil, err
			}
			report[j] = level
		}

		reports[i] = report
	}

	return reports, nil
}

func isReportSafe(report []int) bool {
	increasing := true
	decreasing := true
	largestDifference := 0.0

	for i := range report {
		if i + 1 >= len(report) {
			break
		}

		if report[i + 1] > report[i] {
			decreasing = false
		} else if report[i + 1] < report[i] {
			increasing = false
		} else {
			decreasing = false
			increasing = false
		}

		difference := math.Abs(float64(report[i + 1]) - float64(report[i]))
		if difference > largestDifference {
			largestDifference = difference
		}
	}

	return (increasing || decreasing) && largestDifference <= 3.0
}


func Day2Part1(puzzleInputName string) (string, error) {
	reports, err := parseReportsFromPuzzleInput(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}

	numSafe := len(reports)

	for _, report := range reports {
		if !isReportSafe(report) {
			numSafe--
		}
	}
	return strconv.Itoa(numSafe), nil
}

func Day2Part2(puzzleInputName string) (string, error) {
	reports, err := parseReportsFromPuzzleInput(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}

	numSafe := len(reports)

	for _, report := range reports {
		if isReportSafe(report) {
			continue
		}

		safe := false
		for i := range report {
			alteredReport := common.CopySliceWithIndexExcluded(report, i)
			if isReportSafe(alteredReport) {
				safe = true
				break
			}
		}
		if !safe {
			numSafe--
		}
	}
	return strconv.Itoa(numSafe), nil
}
