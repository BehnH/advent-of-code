package main

import (
	"fmt"
	"iter"

	"github.com/behnh/advent-of-code/pkg/utils"
)

func main() {
	answer1 := partOne()
	answer2 := partTwo()

	fmt.Println("Part 1:", answer1)
	fmt.Println("Part 2:", answer2)
}

func partOne() int {
	wordSearch, err := getWordSearch()
	if err != nil {
		fmt.Println(fmt.Errorf("invalid input: %w", err))
		return 0
	}

	count := 0
	count += searchForWord("XMAS", wordSearch)

	return count
}

func partTwo() int {
	wordSearch, err := getWordSearch()
	if err != nil {
		fmt.Println(fmt.Errorf("invalid input: %w", err))
		return 0
	}

	count := 0
	// Honestly fuck whoever did this
	count += findWordAsCross("MAS", wordSearch)

	return count
}

type coords struct {
	x, y int
}

var (
	up    = coords{-1, 0}
	down  = coords{1, 0}
	left  = coords{0, -1}
	right = coords{0, 1}
)

func (v coords) plus(w coords) coords {
	return coords{
		x: v.x + w.x,
		y: v.y + w.y,
	}
}

func searchForWord(word string, wordSearch [][]byte) int {
	rowCount := searchRowsForWord(word, wordSearch)
	columnCount := searchColumnsForWord(word, wordSearch)
	diagonalCount := searchDiagonalsForWord(word, wordSearch)

	return rowCount + columnCount + diagonalCount
}

func itr(wordSearch [][]byte, start coords, direction coords) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		pos := start
		for {
			if pos.x < 0 || pos.x >= len(wordSearch) || pos.y < 0 || pos.y >= len(wordSearch[pos.x]) {
				return
			}
			if !yield(wordSearch[pos.x][pos.y]) {
				return
			}
			pos = pos.plus(direction)
		}
	}
}

func searchSeqForWord(seq iter.Seq[byte], word string) int {
	want := 0
	count := 0

	for letter := range seq {
		switch {
		case letter == word[0]:
			want = 1
		case letter == word[want]:
			want++
			if want == len(word) {
				count++
				want = 0
			}
		default:
			want = 0
		}
	}

	return count
}

func searchRowsForWord(word string, wordSearch [][]byte) int {
	count := 0

	for row := range wordSearch {
		count += searchSeqForWord(itr(wordSearch, coords{row, 0}, right), word)
		count += searchSeqForWord(itr(wordSearch, coords{row, len(wordSearch[row]) - 1}, left), word)
	}

	fmt.Println(count)

	return count
}

func searchColumnsForWord(word string, wordSearch [][]byte) int {
	count := 0

	for col := range wordSearch[0] {
		count += searchSeqForWord(itr(wordSearch, coords{0, col}, down), word)
		count += searchSeqForWord(itr(wordSearch, coords{len(wordSearch[0]) - 1, col}, up), word)
	}

	return count
}

func searchDiagonalsForWord(word string, wordSearch [][]byte) int {
	count := 0

	directions := []coords{
		up.plus(left),
		up.plus(right),
		down.plus(left),
		down.plus(right),
	}

	for _, dir := range directions {
		for row := range wordSearch {
			count += searchSeqForWord(itr(wordSearch, coords{row, 0}, dir), word)
			count += searchSeqForWord(itr(wordSearch, coords{row, len(wordSearch[row]) - 1}, dir), word)
		}

		for col := range wordSearch[0] {
			// Skip starting points already covered by previous loop.
			if col == 0 || col == len(wordSearch[0])-1 {
				continue
			}

			count += searchSeqForWord(itr(wordSearch, coords{0, col}, dir), word)
			count += searchSeqForWord(itr(wordSearch, coords{len(wordSearch[0]) - 1, col}, dir), word)
		}
	}

	return count
}

func findWordAsCross(word string, wordSearch [][]byte) int {
	count := 0

	for row := range wordSearch {
		for col := range wordSearch[row] {
			if wordSearch[row][col] == word[1] {
				if is3LetterWordAsCross(wordSearch, row, col, word) {
					count++
				}
			}
		}
	}

	return count
}

func is3LetterWordAsCross(wordsearch [][]byte, x, y int, word string) bool {
	if x < 1 || x >= len(wordsearch)-1 || y < 1 || y >= len(wordsearch[x])-1 {
		return false
	}

	tL := wordsearch[x-1][y-1]
	tR := wordsearch[x-1][y+1]
	bL := wordsearch[x+1][y-1]
	bR := wordsearch[x+1][y+1]

	diag1 := (tL == word[0] && bR == word[2]) || (tL == word[2] && bR == word[0])
	diag2 := (tR == word[0] && bL == word[2]) || (tR == word[2] && bL == word[0])

	return diag1 && diag2
}

func getWordSearch() ([][]byte, error) {
	lines, err := utils.GetLines("input.txt")
	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("not a valid word search")
	}

	rowLen := len(lines[0])
	for _, line := range lines {
		if len(line) != rowLen {
			return nil, fmt.Errorf("not a rectangular word search")
		}
	}

	wordSearch := make([][]byte, len(lines))
	for row, line := range lines {
		wordSearch[row] = []byte(line)
	}

	return wordSearch, nil
}
