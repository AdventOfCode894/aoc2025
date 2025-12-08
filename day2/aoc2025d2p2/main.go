package main

import (
	"io"
	"strconv"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)
	tr := pr.LineTokenReader()
	invalidSum := 0
	for {
		start, ok := tr.NextUint('-', 10)
		if !ok {
			break
		}
		end, ok := tr.NextUint(',', 10)
		if !ok {
			break
		}
		for i := start; i <= end; i++ {
			if isInvalidID(i) {
				invalidSum += int(i)
			}
		}
	}

	return invalidSum, nil
}

func isInvalidID(id uint) bool {
	s := strconv.FormatUint(uint64(id), 10)
	n := len(s)
	for digits := 1; digits <= n/2; digits++ {
		if n%digits != 0 {
			continue
		}
		if isIDRepetition(s, digits) {
			return true
		}
	}

	return false
}

func isIDRepetition(s string, digits int) bool {
	for r := s[digits:]; len(r) > 0; r = r[digits:] {
		if s[:digits] != r[:digits] {
			return false
		}
	}

	return true
}
