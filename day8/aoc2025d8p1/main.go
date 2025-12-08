package main

import (
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

	connections := make(map[int][]int)
	for i := range 1000 {
		connections[potentialWires[i].idx1] = append(connections[potentialWires[i].idx1], potentialWires[i].idx2)
		connections[potentialWires[i].idx2] = append(connections[potentialWires[i].idx2], potentialWires[i].idx1)
	}

	var cliqueSizes []int
	seen := make(map[int]struct{})
	for i := range points {
		if _, ok := seen[i]; ok {
			continue
		}
		subSeen := make(map[int]struct{})
		walkClique(i, connections, subSeen)
		cliqueSizes = append(cliqueSizes, len(subSeen))
		for j := range subSeen {
			seen[j] = struct{}{}
		}
	}

	slices.Sort(cliqueSizes)

	return cliqueSizes[len(cliqueSizes)-1] * cliqueSizes[len(cliqueSizes)-2] * cliqueSizes[len(cliqueSizes)-3], nil
}

func walkClique(i int, connections map[int][]int, seen map[int]struct{}) {
	if _, ok := seen[i]; ok {
		return
	}
	seen[i] = struct{}{}
	for _, next := range connections[i] {
		walkClique(next, connections, seen)
	}
}
