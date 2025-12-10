package main

import (
	"io"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type point struct {
	x int
	y int
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var redTiles []point
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		x, _ := tr.NextInt(',', 10)
		y, _ := tr.NextInt(aocio.EOLDelim, 10)
		redTiles = append(redTiles, point{x: x, y: y})
	}

	maxArea := 0
	for i, tile1 := range redTiles {
		for _, tile2 := range redTiles[i+1:] {
			w := tile1.x - tile2.x
			if w < 0 {
				w = -w
			}
			w++
			h := tile1.y - tile2.y
			if h < 0 {
				h = -h
			}
			h++
			area := w * h
			maxArea = max(maxArea, area)
		}
	}

	return maxArea, nil
}
