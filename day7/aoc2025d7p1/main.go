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

	var tachyonEntries []bool
	splits := 0
	for pr.NextLine() {
		manifoldLine := pr.LineRunes()
		if len(tachyonEntries) < 1 {
			tachyonEntries = make([]bool, len(manifoldLine))
			for i, b := range manifoldLine {
				if b == 'S' {
					tachyonEntries[i] = true
				}
			}
			continue
		}
		for i, b := range manifoldLine {
			if b != '^' {
				continue
			}
			if tachyonEntries[i] {
				if !tachyonEntries[i-1] {
					tachyonEntries[i-1] = true
				}
				if !tachyonEntries[i+1] {
					tachyonEntries[i+1] = true
				}
				tachyonEntries[i] = false
				splits++
			}
		}
	}

	return splits, nil
}
