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

	// Deinfe the nr of matches variable
	nrOfMatches := 0

	for sc.Scan() {
		// Get the current line
		line := sc.Text()

		// Define the variables
		var n1, n2, n3, n4 int

		fmt.Sscanf(line, "%d-%d,%d-%d", &n1, &n2, &n3, &n4)
		if (n1-n3)*(n2-n4) <= 0 {
			nrOfMatches++
		}
	}

	fmt.Println(nrOfMatches)
}
