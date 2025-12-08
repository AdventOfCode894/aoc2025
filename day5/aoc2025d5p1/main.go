package main

import (
	"io"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type idRange struct {
	start int
	end   int
}

func (ir idRange) InRange(id int) bool {
	return ir.start <= id && id <= ir.end
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var freshRanges []idRange
	for pr.NextLine() && !pr.IsLineEmpty() {
		tr := pr.LineTokenReader()
		start, _ := tr.NextInt('-', 10)
		end, _ := tr.NextInt(aocio.EOLDelim, 10)
		freshRanges = append(freshRanges, idRange{start: start, end: end})
	}

	numFresh := 0
	for pr.NextLine() {
		id := pr.ReadIntLine(10)
		for _, ir := range freshRanges {
			if ir.InRange(id) {
				numFresh++
				break
			}
		}
	}

	return numFresh, nil
}
