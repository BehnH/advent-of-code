package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/behnh/advent-of-code/pkg/utils"
)

func main() {
	answer1, invalidElems := partOne()
	answer2 := partTwo(invalidElems)

	fmt.Println("Part 1:", answer1)
	fmt.Println("Part 2:", answer2)
}

func partOne() (int, [][]int) {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return 0, nil
	}

	total := 0
	rules, updates := splitLines(lines)
	r := parseRules(rules)

	var valid, invalid [][]int

	for _, update := range updates {
		isValid := true
		for i, v := range update {
			if e, ok := r[v]; ok {
				if i != len(update)-1 {
					for _, a := range update[i+1:] {
						if !slices.Contains(e.after, a) {
							isValid = false
						}
					}
				}
				if i != 0 {
					for _, b := range update[:i] {
						if !slices.Contains(e.before, b) {
							isValid = false
						}
					}
				}
			}
		}
		if isValid {
			valid = append(valid, update)
		} else {
			invalid = append(invalid, update)
		}
	}

	total = sumMiddleElems(valid)
	return total, invalid
}

func partTwo(invalidElems [][]int) int {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return 0
	}

	total := 0
	rules, _ := splitLines(lines)
	r := parseRules(rules)

	for _, elem := range invalidElems {
		var fixed []int
		for _, v := range elem {
			if len(fixed) == 0 {
				fixed = append(fixed, v)
				continue
			}
			for _, f := range fixed {
				if slices.Contains(r[f].before, v) {
					fixed = append([]int{v}, fixed...)
					break
				}
				if slices.Contains(r[f].after, v) {
					for i := len(fixed) - 1; i >= 0; i-- {
						if slices.Contains(r[fixed[i]].after, v) {
							fixed = append(fixed[:i+1], append([]int{v}, fixed[i+1:]...)...)
							break
						}
					}
					break
				}
			}
		}
		total += sumMiddleElems([][]int{fixed})
	}

	return total
}

func splitLines(lines []string) ([][]int, [][]int) {
	var rules, updates [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			r := strings.Split(line, "|")
			var rule []int
			for _, n := range r {
				rule = append(rule, utils.StrToInt(n))
			}
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			r := strings.Split(line, ",")
			var update []int
			for _, n := range r {
				update = append(update, utils.StrToInt(n))
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

type ruleStruct struct {
	before []int
	after  []int
}

func parseRules(rules [][]int) map[int]ruleStruct {
	r := make(map[int]ruleStruct)

	for _, rule := range rules {
		e, ok := r[rule[0]]
		if !ok {
			var entry ruleStruct
			entry.after = append(entry.after, rule[1])
			r[rule[0]] = entry
		} else {
			e.after = append(e.after, rule[1])
			r[rule[0]] = e
		}

		e, ok = r[rule[1]]
		if !ok {
			var entry ruleStruct
			entry.before = append(entry.before, rule[0])
			r[rule[1]] = entry
		} else {
			e.before = append(e.before, rule[0])
			r[rule[1]] = e
		}
	}

	return r
}

func sumMiddleElems(elems [][]int) int {
	total := 0

	for _, v := range elems {
		// Get the index of the middle element
		m := int(math.Ceil(float64(len(v))/2)) - 1
		total += v[m]
	}

	return total
}
