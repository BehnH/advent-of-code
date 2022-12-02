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
	// Rock is X or A (1), Paper is Y or B (2), Scissors is Z or C (3)
	scores := map[string]struct { f, s int } {
		// Rock vs Rock, Rock vs Paper, Rock vs Scissors
		"A X": { 1 + 3, 1 + 3 }, "A Y": { 1 + 0, 2 + 6 }, "A Z": { 1 + 6, 3 + 0 },
		// Paper vs Rock, Paper vs Paper, Paper vs Scissors
		"B X": { 2 + 6, 1 + 0 }, "B Y": { 2 + 3, 2 + 3 }, "B Z": { 2 + 0, 3 + 6 },
		// Scissors vs Rock, Scissors vs Paper, Scissors vs Scissors
		"C X": { 3 + 0, 1 + 6 }, "C Y": { 3 + 6, 2 + 0 }, "C Z": { 3 + 3, 3 + 3 },
	}

	score, score1 := 0, 0
	for sc.Scan() {
		score += scores[sc.Text()].f
		score1 += scores[sc.Text()].s
	}

	fmt.Println(score, score1)
}