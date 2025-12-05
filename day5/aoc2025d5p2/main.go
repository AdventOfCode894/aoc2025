package main

import (
	"io"

	"github.com/AdventOfCodee894/aoc2025/internal/aocio"
	"github.com/AdventOfCodee894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type idRange struct {
	Start   int
	End     int
	Deleted bool
}

func (ir idRange) InRange(id int) bool {
	return ir.Start <= id && id <= ir.End
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var freshRanges []idRange
	for pr.NextLine() && !pr.IsLineEmpty() {
		tr := pr.LineTokenReader()
		start, _ := tr.NextInt('-', 10)
		end, _ := tr.NextInt(aocio.EOLDelim, 10)

		for _, ir := range freshRanges {
			if ir.Deleted {
				continue
			}
			if ir.InRange(start) {
				start = ir.End + 1
			}
			if ir.InRange(end) {
				end = ir.Start - 1
			}
		}
		if start > end {
			continue
		}

		newRange := idRange{Start: start, End: end}

		for i, ir := range freshRanges {
			if newRange.InRange(ir.Start) && newRange.InRange(ir.End) {
				freshRanges[i].Deleted = true
			}
		}

		freshRanges = append(freshRanges, newRange)
	}

	numFresh := 0
	for _, ir := range freshRanges {
		if ir.Deleted {
			continue
		}
		numFresh += ir.End - ir.Start + 1
	}

	return numFresh, nil
}
