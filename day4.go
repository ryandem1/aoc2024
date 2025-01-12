package main

import (
	"github.com/ryandem1/aoc2024/common"
	"slices"
	"sort"
	"strconv"
)

func Day4Part1(puzzleInputName string) (string, error) {
	puzzleInputLines, err := common.GetPuzzleInputByLines(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}
	targetWord := "XMAS"
	targetWordLen := len(targetWord)
	targetWordCount := 0

	for i, line := range puzzleInputLines {
		for j := range line {
			if line[j] == 'X' {
				wordEnd := targetWordLen - 1
				rightUnbound := j + wordEnd < len(line)
				leftUnbound := j - wordEnd >= 0
				bottomUnbound := i + wordEnd < len(puzzleInputLines)
				topUnbound := i - wordEnd >= 0

				var word string
				wordChars := make([]uint8, targetWordLen)

				if rightUnbound {
					word = line[j : j+targetWordLen]

					if word == targetWord {
						targetWordCount++
					}

					wordChars = make([]uint8, targetWordLen)
				}

				if rightUnbound && bottomUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i+q][j+q]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}

					wordChars = make([]uint8, targetWordLen)
				}

				if bottomUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i+q][j]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}

				if bottomUnbound && leftUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i+q][j-q]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}

				if leftUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i][j-q]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}

				if leftUnbound && topUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i-q][j-q]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}

				if topUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i-q][j]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}

				if topUnbound && rightUnbound {
					for q := range targetWordLen {
						wordChars[q] = puzzleInputLines[i-q][j+q]
					}
					word = string(wordChars)

					if word == targetWord {
						targetWordCount++
					}
					wordChars = make([]uint8, targetWordLen)
				}
			}
		}
	}
	return strconv.Itoa(targetWordCount), nil
}


func Day4Part2(puzzleInputName string) (string, error) {
	puzzleInputLines, err := common.GetPuzzleInputByLines(puzzleInputName)
	if err != nil {
		return stringErrorValue, err
	}
	numXmas := 0
	crossLetters := []string{"M", "S"}  // must be only 2 letters

	for i, line := range puzzleInputLines {
		for j := range line {
			if line[j] == 'A' {
				rightUnbound := j + 1 < len(line)
				leftUnbound := j - 1 >= 0
				bottomUnbound := i + 1 < len(puzzleInputLines)
				topUnbound := i - 1 >= 0

				if !(topUnbound && bottomUnbound && leftUnbound && rightUnbound) {
					continue
				}
				topRightChar := puzzleInputLines[i-1][j+1]
				topLeftChar := puzzleInputLines[i-1][j-1]
				bottomRightChar := puzzleInputLines[i+1][j+1]
				bottomLeftChar := puzzleInputLines[i+1][j-1]

				cross1 := []string{string(bottomLeftChar), string(topRightChar)}
				cross2 := []string{string(topLeftChar), string(bottomRightChar)}

				sort.Strings(cross1)
				sort.Strings(cross2)

				if slices.Equal(cross1, crossLetters) && slices.Equal(cross2, crossLetters) {
					numXmas++
				}
			}
		}
	}

	return strconv.Itoa(numXmas), nil
}
