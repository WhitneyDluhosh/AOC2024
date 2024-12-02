package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func absoluteDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func sumOfDifference(left []int, right []int) int {
	var sum = 0
	for i := 0; i < len(left); i++ {
		diff := absoluteDiff(left[i], right[i])
		sum = sum + diff
		println(left[i], right[i], diff)
	}
	return sum
}

func sumOfSimilarity(left []int, occurrencesInRight map[int]int) int {
	var sum = 0
	for i := 0; i < len(left); i++ {
		val := left[i] * occurrencesInRight[left[i]]
		sum = sum + val
		println(left[i], occurrencesInRight[left[i]], val)
	}
	return sum
}

func main() {
	// parse the input into two lists 1st entry - 1st list 2nd entry -2nd list
	file, _ := os.Open("test_input.txt")
	defer file.Close()

	r := bufio.NewReader(file)
	var left []int
	var right []int
	occurrenceTracker := make(map[int]int)
	// parse lists from input file
	for {

		line, _, err := r.ReadLine()
		if len(line) > 0 {
			entries := strings.Fields(string(line))
			leftEntry, _ := strconv.Atoi(entries[0])
			rightEntry, _ := strconv.Atoi(entries[1])
			i, ok := occurrenceTracker[rightEntry]
			if ok {
				occurrenceTracker[rightEntry] = i + 1
			} else {
				occurrenceTracker[rightEntry] = 1
			}
			left = append(left, leftEntry)
			right = append(right, rightEntry)
		}
		if err != nil {
			break
		}
	}
	// sort the values in the list (asc/dsc doesn't matter due to the comparison)
	slices.Sort(left)
	slices.Sort(right)
	// compare the two lists tracking the difference between each set of values and tallying sum
	println("Difference: ", sumOfDifference(left, right))
	println("Similarity: ", sumOfSimilarity(left, occurrenceTracker))

}

// Second Problem
// get the sum of the similarity between left and right
// similarity is the number of occurrences of a value in the left list in the right times the number
// when we intake the right list we want to make a hashmap tallying the numbers to their values which we can increase if we encounter

// First Problem
// pair up the smallest number in the left list
// with the smallest number in the right list
// then the second with the second in the right and so on
// with each pair figure ou t how far apart the two numbers are
// then add up all those distances
