package main

import (
	"fmt"
	"io"

	"github.com/AdventOfCodee894/aoc2025/internal/aocio"
	"github.com/AdventOfCodee894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	dial := 50
	zeroes := 0
	for pr.NextNonEmptyLine() {
		tr := pr.LineTokenReader()

		dir, ok := tr.NextRune()
		if !ok {
			return 0, fmt.Errorf("could not read dir")
		}

		amount, ok := tr.NextInt(aocio.EOLDelim, 10)
		if !ok {
			return 0, fmt.Errorf("could not read amount")
		}

		switch dir {
		case 'L':
			dial = (dial - amount) % 100
		case 'R':
			dial = (dial + amount) % 100
		default:
			return 0, fmt.Errorf("unexpected dir: %v", dir)
		}

		if dial == 0 {
			zeroes++
		}
	}

	return zeroes, nil
}
