package main

import (
	"io"
	"strconv"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var blackboard [][]rune
	maxLen := 0
	for pr.NextLine() {
		blackboard = append(blackboard, pr.LineRunes())
		maxLen = max(maxLen, len(pr.LineRunes()))
	}
	n := len(blackboard) - 1

	var op rune
	var operands []int
	answer := 0
	for i := maxLen - 1; i >= 0; i-- {
		var num []rune
		for _, line := range blackboard[:n] {
			if len(line) > i && line[i] != ' ' {
				num = append(num, line[i])
			}
		}
		if len(blackboard[n]) > i && blackboard[n][i] != ' ' {
			op = blackboard[n][i]
		}
		if len(num) > 0 {
			x, _ := strconv.Atoi(string(num))
			operands = append(operands, x)
		}
		if len(num) < 1 || i == 0 {
			answer += doMath(op, operands)
			operands = operands[:0]
		}
	}

	return answer, nil
}

func doMath(op rune, operands []int) int {
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
