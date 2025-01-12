package main

import (
	"flag"
	"log"
	"time"
)

const (
	stringErrorValue string = ""
)

func main() {
	puzzleFlag := flag.String(
		"puzzle",
		"day1part1",
		"which puzzle to run",
	)

	inputFlag := flag.String(
		"input",
		"day1part1example",
		"name of puzzle input file to use.",
	)
	flag.Parse()

	puzzle, ok := map[string]func(string)(string, error) {
		"day1part1": Day1Part1,
		"day1part2": Day1Part2,
		"day2part1": Day2Part1,
		"day2part2": Day2Part2,
		"day3part1": Day3Part1,
		"day3part2": Day3Part2,
		"day4part1": Day4Part1,
		"day4part2": Day4Part2,
		"day5part1": Day5Part1,
		"day5part2": Day5Part2,
	}[*puzzleFlag]

	if !ok {
		log.Fatalf("invalid puzzle: '%s'", *puzzleFlag)
	}

	now := time.Now()
	answer, err := puzzle(*inputFlag)
	elapsed := time.Since(now)

	if err != nil {
		panic(err)
	}
	log.Printf(
		"Elapsed: %s Answer for '%s' with input file '%s' is: '%s'\n",
		elapsed,
		*puzzleFlag,
		*inputFlag,
		answer,
	)
}
