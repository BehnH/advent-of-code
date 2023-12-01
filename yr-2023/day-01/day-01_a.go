package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input
	lines, err := parse("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println(partOne(lines))
}

func partOne(file string) int {
	// Get numbers from each line
	r := regexp.MustCompile(`[^\d\n]+`)
	file = r.ReplaceAllString(file, "")
	lines := strings.Split(file, "\n")
	sum := 0
	for _, line := range lines {
		first := line[0]
		last := line[len(line)-1]
		// Add first and last number
		f, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		sum += f
	}
	return sum
}

func parse(file string) (string, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
