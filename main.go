package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	PuzzleInputsFolder string = "inputs" // relative to the CWD
	stringErrorValue string = ""
)


// GetPuzzleInput expects inputs to be stored in an `ìnputs` folder
// within the current working directory.
// Due to licensing, my inputs are not stored in the project.
func GetPuzzleInput(fileName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return stringErrorValue, err
	}
	puzzleInputPath := filepath.Join(cwd, PuzzleInputsFolder, fileName)

	file, err := os.Open(puzzleInputPath)
	if err != nil {
		return stringErrorValue, err
	}

	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	fileContent := ""
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}

	return fileContent, nil
}

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
