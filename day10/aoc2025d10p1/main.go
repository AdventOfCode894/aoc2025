package main

import (
	"errors"
	"io"
	"strings"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
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
	lightsDesired, _ := tr.NextString(']')

	var buttons [][]bool
	for {
		tr.NextRune()
		rn, _ := tr.NextRune()
		if rn != '(' {
			break
		}
		buttonLights := tr.NextIntArray(',', ')', 10)
		button := make([]bool, len(lightsDesired))
		for _, l := range buttonLights {
			button[l] = true
		}
		buttons = append(buttons, button)
	}

	start := strings.Repeat(".", len(lightsDesired))
	distances := make(map[string]int)
	distances[start] = 0
	queue := []string{start}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		nextDist := distances[state] + 1
		for _, b := range buttons {
			next := []rune(state)
			for l, toggle := range b {
				if !toggle {
					continue
				}
				if next[l] == '#' {
					next[l] = '.'
				} else {
					next[l] = '#'
				}
			}
			existingDist, ok := distances[string(next)]
			if ok && existingDist <= nextDist {
				continue
			}
			if string(next) == lightsDesired {
				return nextDist, nil
			}
			distances[string(next)] = nextDist
			queue = append(queue, string(next))
		}
	}

	return 0, errors.New("solution unreachable")
}
