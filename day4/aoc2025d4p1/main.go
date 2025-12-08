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

	var grid [][]bool
	for pr.NextNonEmptyLine() {
		rn := pr.LineRunes()
		row := make([]bool, len(rn))
		for i, v := range rn {
			row[i] = v == '@'
		}
		grid = append(grid, row)
	}

	accessible := 0
	for y, row := range grid {
		for x, roll := range row {
			if !roll {
				continue
			}
			surrounding := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					x2 := x + i
					y2 := y + j
					if y2 < 0 || y2 >= len(grid) {
						continue
					}
					if x2 < 0 || x2 >= len(grid[y2]) {
						continue
					}
					if grid[y2][x2] {
						surrounding++
					}
				}
			}
			if surrounding < 4 {
				accessible++
			}
		}
	}

	return accessible, nil
}
