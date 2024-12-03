package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

	total := 0

	for _, line := range lines {
		vals := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)").FindAllString(line, -1)

		for _, v := range vals {
			// Pick out the numbers
			n1, n2 := mulToInts(v)

			t := n1 * n2
			total += t
		}
	}

	return total
}

func partTwo() int {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return 0
	}

	total := 0
	r := regexp.MustCompile("((don't[(][)])|(do[(][)])|(mul[(]([[:digit:]]{1,3}),([[:digit:]]{1,3})[)]))")

	toDo := true
	for _, line := range lines {
		for _, m := range r.FindAllStringSubmatch(line, -1) {
			if len(m[4]) > 0 {
				if toDo {
					m0 := strToInt(m[5])
					m1 := strToInt(m[6])

					total += m0 * m1
				}
			} else if len(m[3]) > 0 {
				toDo = true
			} else if len(m[2]) > 0 {
				toDo = false
			}
		}
	}
	return total
}

func mulToInts(s string) (int, int) {
	n := regexp.MustCompile("(\\d{1,3}),(\\d{1,3})").FindString(s)

	lh, _ := strconv.Atoi(strings.Split(n, ",")[0])
	rh, _ := strconv.Atoi(strings.Split(n, ",")[1])

	return lh, rh
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
