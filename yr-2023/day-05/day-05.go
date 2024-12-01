package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	p1, err := partOne(lines)
	fmt.Println("Part 1:", p1)

	// Part 2
	p2, err := partTwo(lines)
	fmt.Println("Part 2:", p2)
}

func partOne(lines []string) (int, error) {
	a, err := almanacFromLines(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse almanac: %w", err)
	}

	lowestLoc := math.MaxInt
	for _, seed := range a.seeds {
		seed = a.convert(seed)
		if seed < lowestLoc {
			lowestLoc = seed
		}
	}
	return lowestLoc, nil
}

func partTwo(lines []string) (int, error) {
	a, err := almanacFromLines(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse almanac: %w", err)
	}

	seedRangeLen := len(a.seeds) / 2
	// Wooo goroutines
	lowestLocs := make(chan int, seedRangeLen)
	for i := 0; i < seedRangeLen; i++ {
		go func(i int) {
			start := a.seeds[i*2]
			fin := start + a.seeds[i*2+1]
			log.Printf("worker %d: start=%d, fin=%d", i, start, fin)

			lowestLoc := math.MaxInt
			for seed := start; seed < fin; seed++ {
				loc := a.convert(seed)
				if loc < lowestLoc {
					lowestLoc = loc
				}
			}

			log.Printf("worker %d: lowestLoc=%d", i, lowestLoc)
			lowestLocs <- lowestLoc
		}(i)
	}

	lowestLoc := math.MaxInt
	for i := 0; i < seedRangeLen; i++ {
		loc := <-lowestLocs
		if loc < lowestLoc {
			lowestLoc = loc
		}
	}

	_, err = fmt.Printf("lowestLoc=%d\n", lowestLoc)
	if err != nil {
		return 0, fmt.Errorf("failed to print: %w", err)
	}
	return lowestLoc, nil
}

var categories = [...]string{
	"seed",
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

type almanacRange struct {
	destinationStart int
	sourceStart      int
	length           int
}

type almanacMap struct {
	sourceCategory      string
	destinationCategory string
	ranges              []almanacRange
}

type almanac struct {
	seeds []int
	maps  []almanacMap
}

func (m almanacMap) convert(n int) int {
	for _, r := range m.ranges {
		if n >= r.sourceStart && n < r.sourceStart+r.length {
			return r.destinationStart + (n - r.sourceStart)
		}
	}
	return n
}

func (a almanac) convert(n int) int {
	for _, m := range a.maps {
		n = m.convert(n)
	}
	return n
}

func almanacFromLines(lines []string) (*almanac, error) {
	chunks := splitSlice(lines, func(s string) bool {
		return s == ""
	})
	// Quick sanity check to make sure the chunks match the num of categories
	if len(chunks) != len(categories) {
		panic("chunks and categories don't match")
	}
	// Then check that the first chunk only has one line of seeds
	if len(chunks[0]) != 1 {
		panic("first chunk should only have one line")
	}

	seeds, err := seedsFromString(chunks[0][0])
	if err != nil {
		panic(err)
	}
	if len(seeds)%2 != 0 {
		panic("invalid number of seeds")
	}

	var maps []almanacMap
	for i, chunk := range chunks[1:] {
		m, err := almanacMapFromChunk(chunk)
		if err != nil {
			panic(err)
		}
		if m.sourceCategory != categories[i] {
			panic("invalid source category")
		}
		if strings.Trim(m.destinationCategory, " ") != strings.ToLower(categories[i+1]) {
			fmt.Println(m.destinationCategory, categories[i+1])
			panic("invalid destination category")
		}
		maps = append(maps, m)
	}

	return &almanac{
		seeds: seeds,
		maps:  maps,
	}, nil
}

func almanacMapFromChunk(chunk []string) (almanacMap, error) {
	source, dest, err := categoriesFromHeader(chunk[0])
	if err != nil {
		return almanacMap{}, fmt.Errorf("failed to parse header: %w", err)
	}

	var ranges []almanacRange
	for _, line := range chunk[1:] {
		nums, err := integersFromString(line, " ")
		if err != nil {
			return almanacMap{}, fmt.Errorf("failed to parse line: %w", err)
		}
		if len(nums) != 3 {
			return almanacMap{}, fmt.Errorf("invalid number of integers")
		}
		ranges = append(ranges, almanacRange{
			destinationStart: nums[0],
			sourceStart:      nums[1],
			length:           nums[2],
		})
	}

	return almanacMap{
		sourceCategory:      source,
		destinationCategory: dest,
		ranges:              ranges,
	}, nil
}

func categoriesFromHeader(header string) (string, string, error) {
	header = strings.TrimSuffix(header, "map:")
	parts := strings.Split(header, "-to-")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid header")
	}
	return parts[0], parts[1], nil
}

func seedsFromString(s string) ([]int, error) {
	s = strings.TrimPrefix(s, "seeds: ")
	return integersFromString(s, " ")
}

func splitSlice(s []string, fn func(string) bool) [][]string {
	var chunks [][]string
	var chunk []string

	for _, v := range s {
		if fn(v) {
			chunks = append(chunks, chunk)
			chunk = nil
		} else {
			chunk = append(chunk, v)
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}
	return chunks
}

func integersFromString(s, sep string) ([]int, error) {
	words := strings.Split(s, sep)
	nums := make([]int, len(words))
	for i, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

func parse(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		panic(s.Err())
	}

	return lines, nil
}
