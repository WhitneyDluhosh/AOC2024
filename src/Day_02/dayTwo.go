package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isSafe(puzzle []string) bool {
	// puzzle length should be 5 but just incase
	var increasing = false
	var decreasing = false
	for i := 0; i < len(puzzle)-1; i++ {
		left, _ := strconv.Atoi(puzzle[i])
		right, _ := strconv.Atoi(puzzle[i+1])
		var diff = 1
		if left < right {
			increasing = true
			diff = right - left
		} else {
			decreasing = true
			diff = left - right
		}
		// exit early if unsafe
		if diff < 1 || diff > 3 {
			return false
		}
		if increasing && decreasing {
			return false
		}

	}
	return true

}

func puzzleSliceNCheck(puzzle []string, i int) bool {
	var puzzleLength = len(puzzle)
	var slicedPuzzle []string
	cpy := make([]string, len(puzzle))
	copy(cpy, puzzle)
	if i == 0 {
		slicedPuzzle = append([]string{}, cpy[1:puzzleLength]...)
	} else if i == puzzleLength-1 {
		slicedPuzzle = append([]string{}, cpy[0:puzzleLength-1]...)
	} else {
		anotherSlice := append([]string{}, cpy[i+1:puzzleLength]...)
		slicedPuzzle = append(cpy[0:i], anotherSlice...)
	}
	return isSafe(slicedPuzzle)

}

func isSafeDampened(puzzle []string) bool {
	// puzzle length should be 5 but just incase
	for j := 0; j < len(puzzle); j++ {
		if puzzleSliceNCheck(puzzle, j) {
			return true
		}

	}
	return false

}

func main() {
	// parse the input into two lists 1st entry - 1st list 2nd entry -2nd list
	file, _ := os.Open("test_input.txt")
	defer file.Close()
	r := bufio.NewReader(file)
	var safeCount = 0
	var safeCountDampened = 0
	// parse lists from input file
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			entries := strings.Fields(string(line))
			if isSafe(entries) {
				safeCount = safeCount + 1
			} else if isSafeDampened(entries) {
				safeCountDampened = safeCountDampened + 1
			}
		}
		if err != nil {
			break
		}

	}
	println("number of safe: ", safeCount)
	println("number of safeDampened: ", safeCountDampened)
	println("sum: ", safeCount+safeCountDampened)

}
