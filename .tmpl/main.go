package main

import (
	"fmt"

	"github.com/behnh/advent-of-code/pkg/utils"
)

func main() {
	answer1 := partOne()
	answer2 := partTwo()

	fmt.Println("Part 1:", answer1)
	fmt.Println("Part 2:", answer2)
}

func partOne() int {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return 0
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return 0
}

func partTwo() int {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return 0
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return 0
}
