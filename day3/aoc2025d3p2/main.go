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
		jolts += enabledJoltage(pr.ReadUintArrayLine(aocio.NoDelim, 10), 12, 0)
	}

	return int(jolts), nil
}

func enabledJoltage(bank []uint, digits int, accumulation uint) uint {
	if digits < 1 {
		return 0
	}

	pos := maxBatteryPos(bank[:len(bank)-digits+1])
	accumulation = accumulation*10 + bank[pos]
	if digits > 1 {
		accumulation = enabledJoltage(bank[pos+1:], digits-1, accumulation)
	}

	return accumulation
}

func maxBatteryPos(bank []uint) int {
	pos := 0
	for i, b := range bank[1:] {
		if b > bank[pos] {
			pos = i + 1
		}
	}

	return pos
}
