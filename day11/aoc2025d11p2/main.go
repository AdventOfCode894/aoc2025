package main

import (
	"io"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type devState struct {
	processed        bool
	uncheckedInputs  int
	pathsWithNeither int
	pathsWithFFT     int
	pathsWithDAC     int
	pathsWithBoth    int
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	links := make(map[string][]string)
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		from, _ := tr.NextString(':')
		tr.ConsumeSpaces()
		links[from] = tr.NextStringArray(' ', aocio.EOLDelim)
	}

	states := map[string]*devState{"out": new(devState)}
	for from := range links {
		states[from] = new(devState)
	}
	for _, tos := range links {
		for _, to := range tos {
			states[to].uncheckedInputs++
		}
	}
	states["svr"].pathsWithNeither = 1

	for {
		var from string
		for candidate, state := range states {
			if state.processed || state.uncheckedInputs > 0 {
				continue
			}
			from = candidate
			state.processed = true
			break
		}
		if from == "" {
			break
		}

		for _, to := range links[from] {
			states[to].uncheckedInputs--
			switch from {
			case "fft":
				states[to].pathsWithFFT += states[from].pathsWithFFT + states[from].pathsWithNeither
				states[to].pathsWithBoth += states[from].pathsWithBoth + states[from].pathsWithDAC
			case "dac":
				states[to].pathsWithDAC += states[from].pathsWithDAC + states[from].pathsWithNeither
				states[to].pathsWithBoth += states[from].pathsWithBoth + states[from].pathsWithFFT
			default:
				states[to].pathsWithNeither += states[from].pathsWithNeither
				states[to].pathsWithDAC += states[from].pathsWithDAC
				states[to].pathsWithFFT += states[from].pathsWithFFT
				states[to].pathsWithBoth += states[from].pathsWithBoth
			}
		}
	}

	return states["out"].pathsWithBoth, nil
}
