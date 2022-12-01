package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	input, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	// Search for the maximum sum of calories
	maxCalories := 0
	currentElfCals := 0

	for sc.Scan() {
		// Read the lines
		snack, err := strconv.Atoi(sc.Text())
		currentElfCals += snack

		if err != nil {
			if (currentElfCals > maxCalories) {
				// Found a new maximum
				maxCalories = currentElfCals
			}

			// Reset the current elf's calories
			currentElfCals = 0
		}
	}

	// Print the result
	fmt.Println(maxCalories)
}