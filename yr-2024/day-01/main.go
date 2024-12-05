package main

import (
	"fmt"
	"os"
	"sort"
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

	answer, col1, col2 := partOne(lines)
	answer2 := partTwo(col1, col2)

	fmt.Println("Part 1:", answer)
	fmt.Println("Part 2:", answer2)
}

func partOne(input string) (int, []int, []int) {
	var colA []int
	var colB []int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		vA, _ := strconv.Atoi(split[0])
		vB, _ := strconv.Atoi(split[1])

		colA = append(colA, vA)
		colB = append(colB, vB)
	}

	sort.Ints(colA)
	sort.Ints(colB)

	var diffs []int

	for i := 0; i < len(colA); i++ {
		diffs = append(diffs, differ(colA[i], colB[i]))
	}

	var sum int
	for _, diff := range diffs {
		sum += diff
	}

	return sum, colA, colB
}

func partTwo(col1, col2 []int) int {
	counts := make(map[int]int) // map of how manu times a number appears

	// Add all the numbers from col2 to the map with a count of 0
	for i := 0; i < len(col1); i++ {
		counts[col1[i]] = 0
	}

	// Loop through all the numbers in col2 and increment the count in the map
	for i := 0; i < len(col2); i++ {
		if _, ok := counts[col2[i]]; ok {
			counts[col2[i]]++
		}
	}

	// Loop through all the numbers in col1 and multiply them by the count they have in col2
	var sum int
	for i := 0; i < len(col1); i++ {
		sum += col1[i] * counts[col1[i]]
	}

	return sum
}

func differ(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
