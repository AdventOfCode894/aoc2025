package main

import (
	"io"

	"github.com/AdventOfCodee894/aoc2025/internal/aocio"
	"github.com/AdventOfCodee894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)
	jolts := uint(0)
	for pr.NextNonEmptyLine() {
		jolts += enabledJoltage(pr.ReadUintArrayLine(aocio.NoDelim, 10))
	}

	return int(jolts), nil
}

func enabledJoltage(bank []uint) uint {
	first := 0
	for i, b := range bank[1 : len(bank)-1] {
		if b > bank[first] {
			first = 1 + i
		}
	}
	second := first + 1
	for i, b := range bank[first+2:] {
		if b > bank[second] {
			second = first + 2 + i
		}
	}

	return bank[first]*10 + bank[second]
}
