package main

import (
	"io"
	"strconv"

	"github.com/AdventOfCodee894/aoc2025/internal/aocio"
	"github.com/AdventOfCodee894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var equations [][]int
	answer := 0
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		for i := 0; ; i++ {
			tr.ConsumeSpaces()
			b, ok := tr.NextToken(' ')
			if !ok {
				break
			}

			switch string(b) {
			case "*", "+":
				answer += doMath(b[0], equations[i])
			default:
				x, _ := strconv.Atoi(string(b))
				if len(equations) <= i {
					equations = append(equations, nil)
				}
				equations[i] = append(equations[i], x)
			}
		}
	}

	return answer, nil
}

func doMath(op byte, operands []int) int {
	total := operands[0]
	for _, x := range operands[1:] {
		switch op {
		case '+':
			total += x
		case '*':
			total *= x
		}
	}

	return total
}
