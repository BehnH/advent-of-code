package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUnique(s string) bool {
	seen := map[rune]bool{}
	for _, c := range s {
		seen[c] = true
	}
	return len(seen) == len(s)
}

func main()  {
	// Read input
	input, err := os.Open("input.txt")

	// Panic if the input file could not be opened
	if err != nil {
		panic(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Scan()

	chars := sc.Text()

	n1 := 0

	// Write a comment which explains the logic of the following code
	for i := 3; i < len(chars); i++ {
		if i > 12 && isUnique(chars[i-13:i+1]) && n1 == 0 {
			n1 = i + 1
		}
	}

	// Print the result
	fmt.Println(n1)

}