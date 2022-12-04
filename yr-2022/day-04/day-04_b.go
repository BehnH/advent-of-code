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
		var r1, r2, m1 []int

		// Get the numbers from the line
		fmt.Sscanf(line, "%d-%d,%d-%d", &n1, &n2, &n3, &n4)
		// Get the range of the first number set
		for i := n1; i <= n2; i++ {
			r1 = append(r1, i)
		}
		// Get the range of the second number set
		for i := n3; i <= n4; i++ {
			r2 = append(r2, i)
		}

		// Loop through the ranges
		for _, i := range r1 {
			for _, j := range r2 {
				// If the numbers are equal, append them to the m1 slice
				if i == j {
					m1 = append(m1, i)
				}
			}
		}

		// If the length of the m1 slice is greater than 0, increment the nr of matches by 1
		if len(m1) > 0 {
			nrOfMatches++
		}
	}

	// Print the nr of matches
	fmt.Println(nrOfMatches)
}
