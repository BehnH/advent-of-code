package utils

import (
	"fmt"
	"os"
	"slices"
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

func GetLines(filePath string) ([]string, error) {
	lines, err := parse(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading file: %v", err))
		return nil, err
	}

	lineArr := strings.Split(lines, "\n")

	for i, line := range lineArr {
		if line == "" {
			// Remove empty lines
			slices.Delete(lineArr, i, i)
		}
	}

	return lineArr, nil
}

func LinesToIntsArray(lines []string) [][]int {
	var ints [][]int

	for _, line := range lines {
		var numsArray []int
		for _, num := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(num)
			numsArray = append(numsArray, n)
		}
		ints = append(ints, numsArray)
	}

	return ints
}
