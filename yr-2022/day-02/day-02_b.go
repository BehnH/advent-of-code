package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input
	input, err := os.Open("input.txt")

	// Panic if the input file could not be opened
	if err != nil {
		panic(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	// Define the scores
	// Rock is A (1), Paper is B (2), Scissors C (3)
	// X means we have to lose, Y means we have to tie, Z means we have to win
	scores := map[string]struct { f int } {
		// Rock vs Rock, Rock vs Paper, Rock vs Scissors
		"A X": { 0 + 3 }, "A Y": { 3 + 1 }, "A Z": { 6 + 2 },
		// Paper vs Rock, Paper vs Paper, Paper vs Scissors
		"B X": { 0 + 1 }, "B Y": { 3 + 2 }, "B Z": {3 + 6 },
		// Scissors vs Rock, Scissors vs Paper, Scissors vs Scissors
		"C X": { 0 + 2, }, "C Y": { 3 + 3 }, "C Z": { 6 + 1 },
	}

	score := 0
	for sc.Scan() {
		score += scores[sc.Text()].f
	}

	fmt.Println(score)
}