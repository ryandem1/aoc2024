package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const (
	PuzzleInputsFolder string = "inputs" // relative to the CWD
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

// GetPuzzleInputByLines will return the puzzle input as a list of string lines.
// End-of-file and beginning-of-file newlines will NOT be returned.
func GetPuzzleInputByLines(fileName string) ([]string, error) {
	puzzleInput, err  := GetPuzzleInput(fileName)
	if err != nil {
		return nil, err
	}

	puzzleInputLines := strings.Split(
		strings.Trim(puzzleInput, "\n"), "\n",
	)

	return puzzleInputLines, nil
}
