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
	maxCalories1, maxCalories2, maxCalories3 := 0, 0, 0
	currentElfCals := 0

	for sc.Scan() {
		// Read the lines
		snack, err := strconv.Atoi(sc.Text())
		currentElfCals += snack

		if err != nil {
			if (currentElfCals > maxCalories3) {
				// Found a new maximum
				maxCalories3 = currentElfCals
			}

			if (maxCalories3 > maxCalories2) {
				// Found a new maximum
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}

			if (maxCalories2 > maxCalories1) {
				// Found a new maximum
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}

			// Reset the current elf's calories
			currentElfCals = 0
		}
	}

	// Print the result
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}