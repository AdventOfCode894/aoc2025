package main

import (
	"io"
	"math"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
	"github.com/draffensperger/golp"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	answer := 0
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		presses, err := solveMachine(tr)
		if err != nil {
			return 0, err
		}
		answer += presses
	}

	return answer, nil
}

func solveMachine(tr *aocio.TokenReader) (int, error) {
	tr.NextRune()
	tr.NextString(']')

	var buttons [][]int
	for {
		tr.NextRune()
		rn, _ := tr.NextRune()
		if rn != '(' {
			break
		}
		button := tr.NextIntArray(',', ')', 10)
		buttons = append(buttons, button)
	}

	desiredJoltages := tr.NextIntArray(',', '}', 10)

	buttonsForCounter := make([][]int, len(desiredJoltages))
	for i, b := range buttons {
		for _, c := range b {
			buttonsForCounter[c] = append(buttonsForCounter[c], i)
		}
	}

	lp := golp.NewLP(0, len(buttons))
	for c, bs := range buttonsForCounter {
		var entries []golp.Entry
		for _, b := range bs {
			entries = append(entries, golp.Entry{Col: b, Val: 1.0})
		}
		if err := lp.AddConstraintSparse(entries, golp.EQ, float64(desiredJoltages[c])); err != nil {
			return 0, err
		}
	}
	obj := make([]float64, len(buttons))
	for i := range obj {
		obj[i] = 1.0
	}
	lp.SetObjFn(obj)
	for i := range buttons {
		lp.SetInt(i, true)
	}
	lp.Solve()

	minPresses := 0.0
	for _, p := range lp.Variables() {
		minPresses += p
	}

	return int(math.Round(minPresses)), nil
}
