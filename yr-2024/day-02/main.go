package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(file string) (string, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func main() {
	// Read input
	lines, err := parse("input.txt")
	if err != nil {
		panic(err)
	}

	answer1 := partOne(lines)
	answer2 := partTwo(lines)

	fmt.Println("Part 1:", answer1)
	fmt.Println("Part 2:", answer2)
}

func partOne(input string) int {
	safeCount := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var numsArray []int
		nums := strings.Split(line, " ")

		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			numsArray = append(numsArray, n)
		}

		safe := isSafe(numsArray)
		if safe {
			safeCount += 1
		}
	}
	return safeCount
}

func partTwo(input string) int {
	safeCount := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var numsArray []int
		nums := strings.Split(line, " ")

		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			numsArray = append(numsArray, n)
		}

		safe := isSafe(numsArray)
		if safe {
			safeCount += 1
		}

		if !safe {
			removeSafe := checkRemovalSafe(numsArray)
			if removeSafe {
				safeCount += 1
			}
		}
	}
	return safeCount
}

func isSafe(n []int) bool {
	// Loop over the array and ensure the number does not increase or decrease by more than 3 (and at least 1)
	// and that the difference between the numbers is not 0
	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] {
			return false
		}

		if n[i] > n[i+1] {
			if n[i]-n[i+1] > 3 || n[i]-n[i+1] < 1 {
				return false
			}
		} else {
			if n[i+1]-n[i] > 3 || n[i+1]-n[i] < 1 {
				return false
			}
		}
	}

	// Check the array is all increasing - if it is, just return true
	return checkSorted(n)
}

// checkSorted checks if the array is sorted (descending or ascending)
func checkSorted(n []int) bool {
	if slices.IsSorted(n) {
		return true
	} else {
		slices.Reverse(n)
		if slices.IsSorted(n) {
			return true
		}
		return false
	}
}

// checkRemovalSafe checks if removing any element from the array makes it safe
func checkRemovalSafe(n []int) bool {
	for j := range n {
		if isSafe(slices.Concat(n[:j], n[j+1:])) {
			return true
		}
	}

	return false
}
