package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDuplicateVals(strArr []string) string {
	str1, str2, str3 := strArr[0], strArr[1], strArr[2]
	var str string

	for _, str1 := range str1 {
		for _, str2 := range str2 {
			for _, str3 := range str3 {
				if str1 == str2 && str2 == str3 {
					str = string(str1)
				}
			}
		}
	}
	return string(str)
}

func main() {
	// Read input
	input, err := os.Open("input.txt")

	// Panic if the input file could not be opened
	if err != nil {
		panic(err)
	}

	defer input.Close()
	sc := bufio.NewScanner(input)

	result := []string{}

	// Define the scores
	scores := map[string]struct{ f int }{
		"a": {1}, "b": {2}, "c": {3}, "d": {4}, "e": {5}, "f": {6}, "g": {7}, "h": {8},
		"i": {9}, "j": {10}, "k": {11}, "l": {12}, "m": {13}, "n": {14}, "o": {15}, "p": {16},
		"q": {17}, "r": {18}, "s": {19}, "t": {20}, "u": {21}, "v": {22}, "w": {23}, "x": {24},
		"y": {25}, "z": {26}, "A": {27}, "B": {28}, "C": {29}, "D": {30}, "E": {31}, "F": {32},
		"G": {33}, "H": {34}, "I": {35}, "J": {36}, "K": {37}, "L": {38}, "M": {39}, "N": {40},
		"O": {41}, "P": {42}, "Q": {43}, "R": {44}, "S": {45}, "T": {46}, "U": {47}, "V": {48},
		"W": {49}, "X": {50}, "Y": {51}, "Z": {52},
	}

	for sc.Scan() {
		result = append(result, sc.Text())
	}

	// Get the array length
	arrLen := len(result) / 3
	// Create a new array for the badges
	badges := make([]string, arrLen)
	// Loop through the array
	for i := 0; i < arrLen; i++ {
		badges = append(badges, getDuplicateVals(result[i*3 : i*3+3]))
	}

	badgeTotal := 0
	for _, badge := range badges {
		badgeTotal += scores[badge].f
	}

	fmt.Println(badgeTotal)
}
