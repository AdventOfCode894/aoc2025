package main

import (
	"io"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	for range 6 {
		for range 5 {
			pr.NextLine()
		}
	}

	sizes := []int{7, 7, 7, 5, 6, 7}

	treesThatFit := 0
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		w, _ := tr.NextInt('x', 10)
		h, _ := tr.NextInt(':', 10)
		presentCounts := tr.NextIntArray(' ', aocio.EOLDelim, 10)
		totalSpace := 0
		for i := range presentCounts {
			totalSpace += presentCounts[i] * sizes[i]
		}
		if totalSpace < w*h {
			treesThatFit++
		}
	}

	return treesThatFit, nil
}
