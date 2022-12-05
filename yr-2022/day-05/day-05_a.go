package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
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

	/*
	    	        [G] [W]         [Q]
		[Z]         [Q] [M]     [J] [F]
		[V]         [V] [S] [F] [N] [R]
		[T]         [F] [C] [H] [F] [W] [P]
		[B] [L]     [L] [J] [C] [V] [D] [V]
		[J] [V] [F] [N] [T] [T] [C] [Z] [W]
		[G] [R] [Q] [H] [Q] [W] [Z] [G] [B]
		[R] [J] [S] [Z] [R] [S] [D] [L] [J]
		 1   2   3   4   5   6   7   8   9
 	*/

	// Define the initial layout
	stack := [9]Stack{
		{"R", "G", "J", "B", "T", "V", "Z"},
		{"J", "R", "V", "L"},
		{"S", "Q", "F"},
		{"Z", "H", "N", "L", "F", "V", "Q", "G"},
		{"R", "Q", "T", "J", "C", "S", "M", "W"},
		{"S", "W", "T", "C", "H", "F"},
		{"D", "Z", "C", "V", "F", "N", "J"},
		{"L", "G", "Z", "D", "W", "R", "F", "Q"},
		{"J", "B", "W", "V", "P"},
	}

	for sc.Scan() {
		line := sc.Text()
		splitInstructions := strings.Split(line, " ")

		move, _ := strconv.Atoi(splitInstructions[1])
		from, _ := strconv.Atoi(splitInstructions[3])
		to, _ := strconv.Atoi(splitInstructions[5])

		for i := 0; i < move; i++ {
			stack[to-1].Push(stack[from-1].Pop())
		}
	}

	for i := range stack {
		fmt.Printf(stack[i].Pop())
	}

}