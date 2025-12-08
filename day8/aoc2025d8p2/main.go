package main

import (
	"errors"
	"io"
	"slices"

	"github.com/AdventOfCodee894/aoc2025/internal/aocio"
	"github.com/AdventOfCodee894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type point struct {
	x, y, z int
}

func (p point) sqDist(other point) int {
	return (other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y) + (other.z-p.z)*(other.z-p.z)
}

type wire struct {
	idx1   int
	idx2   int
	sqDist int
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var points []point
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		x, _ := tr.NextInt(',', 10)
		y, _ := tr.NextInt(',', 10)
		z, _ := tr.NextInt(aocio.EOLDelim, 10)
		points = append(points, point{x, y, z})
	}

	potentialWires := make([]wire, 0, len(points)*len(points)/2)
	for i, pt1 := range points {
		for j, pt2 := range points[i+1:] {
			potentialWires = append(potentialWires, wire{
				idx1:   i,
				idx2:   i + 1 + j,
				sqDist: pt1.sqDist(pt2),
			})
		}
	}

	slices.SortFunc(potentialWires, func(a, b wire) int {
		if a.sqDist < b.sqDist {
			return -1
		} else if a.sqDist > b.sqDist {
			return 1
		}
		return 0
	})

	cliques := make([]*[]int, len(points))
	for i := range cliques {
		cliques[i] = new([]int)
		*cliques[i] = []int{i}
	}

	for i := range potentialWires {
		idx1 := potentialWires[i].idx1
		idx2 := potentialWires[i].idx2
		if cliques[idx1] == cliques[idx2] {
			continue
		}
		*cliques[idx1] = append(*cliques[idx1], *cliques[idx2]...)
		for _, j := range *cliques[idx2] {
			cliques[j] = cliques[idx1]
		}
		if len(*cliques[idx1]) >= len(points) {
			return points[idx1].x * points[idx2].x, nil
		}
	}

	return 0, errors.New("failed to connect all boxes")
}
