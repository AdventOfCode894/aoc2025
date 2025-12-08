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

	var tachyonWorlds []int
	for pr.NextLine() {
		manifoldLine := pr.LineRunes()
		if len(tachyonWorlds) < 1 {
			tachyonWorlds = make([]int, len(manifoldLine))
			for i, b := range manifoldLine {
				if b == 'S' {
					tachyonWorlds[i] = 1
				}
			}
			continue
		}
		for i, b := range manifoldLine {
			if b != '^' {
				continue
			}
			tachyonWorlds[i-1] += tachyonWorlds[i]
			tachyonWorlds[i+1] += tachyonWorlds[i]
			tachyonWorlds[i] -= tachyonWorlds[i]
		}
	}

	timelines := 0
	for _, worlds := range tachyonWorlds {
		timelines += worlds
	}

	return timelines, nil
}
