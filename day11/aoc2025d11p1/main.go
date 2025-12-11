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

	links := make(map[string][]string)
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		from, _ := tr.NextString(':')
		tr.ConsumeSpaces()
		links[from] = tr.NextStringArray(' ', aocio.EOLDelim)
	}

	return countPathsToOut(links, "you"), nil
}

func countPathsToOut(links map[string][]string, from string) int {
	if from == "out" {
		return 1
	}

	paths := 0
	for _, to := range links[from] {
		paths += countPathsToOut(links, to)
	}

	return paths
}
