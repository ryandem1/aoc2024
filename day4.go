package main

import (
	"github.com/ryandem1/aoc2024/common"
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
	return puzzleInputName, nil
}
