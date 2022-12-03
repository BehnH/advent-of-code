package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeDupelicateVals(str string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range str {
		if _, value := keys[string(entry)]; !value {
			keys[string(entry)] = true
			list = append(list, string(entry))
		}
	}
	return list
}

func getDuplicateVals(str1 string, str2 string) string {
	chars1, chars2 := removeDupelicateVals(str1), removeDupelicateVals(str2)
	var str string

	for _, char1 := range chars1 {
		for _, char2 := range chars2 {
			if char1 == char2 {
				str = char1
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

	dupeTotal := 0

	for sc.Scan() {
		strLen := len(sc.Text()) / 2
		splitStr1, splitStr2 := sc.Text()[:strLen], sc.Text()[strLen:]

		dupeTotal += scores[getDuplicateVals(splitStr1, splitStr2)].f
	}

	fmt.Println(dupeTotal)
}
