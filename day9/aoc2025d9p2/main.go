package main

import (
	"errors"
	"io"
	"slices"

	"github.com/AdventOfCode894/aoc2025/internal/aocio"
	"github.com/AdventOfCode894/aoc2025/internal/aocmain"
)

func main() {
	aocmain.HandlePuzzle(solvePuzzle)
}

type point struct {
	x int
	y int
}

type rectangle struct {
	p1 point
	p2 point
}

func newRectangle(p1 point, p2 point) rectangle {
	return rectangle{
		p1: point{x: min(p1.x, p2.x), y: min(p1.y, p2.y)},
		p2: point{x: max(p1.x, p2.x), y: max(p1.y, p2.y)},
	}
}

func (rect rectangle) Width() int {
	return rect.p2.x - rect.p1.x + 1
}

func (rect rectangle) Height() int {
	return rect.p2.y - rect.p1.y + 1
}

func (rect rectangle) Area() int {
	return rect.Width() * rect.Height()
}

type bound struct {
	minPos   int
	maxPos   int
	orthoPos int
}

func (b bound) Crosses(pos int, orthoMin int, orthoMax int) bool {
	if pos <= b.minPos || pos >= b.maxPos {
		return false
	}
	if b.orthoPos <= orthoMin || b.orthoPos >= orthoMax {
		return false
	}

	return true
}

func (b bound) Touches(pos int, orthoMin int, orthoMax int) bool {
	if pos < b.minPos || pos > b.maxPos {
		return false
	}
	if b.orthoPos < orthoMin || b.orthoPos > orthoMax {
		return false
	}

	return true
}

type tileBoundary struct {
	horizontalBounds []bound
	verticalBounds   []bound
	maxX             int
	maxY             int
}

func newTileBoundary(points []point) tileBoundary {
	tb := tileBoundary{
		horizontalBounds: make([]bound, 0, len(points)),
		verticalBounds:   make([]bound, 0, len(points)),
	}
	for i := range points {
		tb.maxX = max(tb.maxX, points[i].x)
		tb.maxY = max(tb.maxY, points[i].y)

		j := (i + 1) % len(points)
		tile1 := points[i]
		tile2 := points[j]
		if tile1.x == tile2.x {
			tb.verticalBounds = append(tb.verticalBounds, bound{
				minPos:   min(tile1.y, tile2.y),
				maxPos:   max(tile1.y, tile2.y),
				orthoPos: tile1.x,
			})
		} else {
			tb.horizontalBounds = append(tb.horizontalBounds, bound{
				minPos:   min(tile1.x, tile2.x),
				maxPos:   max(tile1.x, tile2.x),
				orthoPos: tile1.y,
			})
		}
	}

	return tb
}

func (tb tileBoundary) Inside(rect rectangle) bool {
	for _, b := range tb.horizontalBounds {
		if b.Crosses(rect.p1.x, rect.p1.y, rect.p2.y) {
			return false
		}
		if b.Crosses(rect.p2.x, rect.p1.y, rect.p2.y) {
			return false
		}
	}
	for _, b := range tb.verticalBounds {
		if b.Crosses(rect.p1.y, rect.p1.x, rect.p2.x) {
			return false
		}
		if b.Crosses(rect.p2.y, rect.p1.x, rect.p2.x) {
			return false
		}
	}
	if tb.isPointOutside(rect.p1.x, 0, rect.p1.x, rect.p1.y, 0, rect.p1.y) {
		return false
	}
	if tb.isPointOutside(rect.p2.x, rect.p2.x, tb.maxX, rect.p1.y, 0, rect.p1.y) {
		return false
	}
	if tb.isPointOutside(rect.p2.x, rect.p2.x, tb.maxX, rect.p2.y, rect.p2.y, tb.maxY) {
		return false
	}
	if tb.isPointOutside(rect.p1.x, 0, rect.p1.x, rect.p2.y, rect.p2.y, tb.maxY) {
		return false
	}

	return true
}

func (tb tileBoundary) isPointOutside(x int, xMin int, xMax int, y int, yMin int, yMax int) bool {
	for _, b := range tb.horizontalBounds {
		if b.Touches(x, yMin, yMax) {
			return false
		}
	}
	for _, b := range tb.verticalBounds {
		if b.Touches(y, xMin, xMax) {
			return false
		}
	}

	return true
}

func solvePuzzle(r io.Reader) (int, error) {
	pr := aocio.NewPuzzleReader(r)

	var redTiles []point
	for pr.NextLine() {
		tr := pr.LineTokenReader()
		x, _ := tr.NextInt(',', 10)
		y, _ := tr.NextInt(aocio.EOLDelim, 10)
		redTiles = append(redTiles, point{x: x, y: y})
	}

	tb := newTileBoundary(redTiles)

	type candidate struct {
		rect rectangle
		area int
	}
	candidates := make([]candidate, 0, len(redTiles)*len(redTiles))
	for i, tile1 := range redTiles {
		for _, tile2 := range redTiles[i+1:] {
			rect := newRectangle(tile1, tile2)
			candidates = append(candidates, candidate{
				rect: rect,
				area: rect.Area(),
			})
		}
	}
	slices.SortFunc(candidates, func(a, b candidate) int {
		if a.area < b.area {
			return 1
		} else if a.area > b.area {
			return -1
		}
		return 0
	})

	for _, c := range candidates {
		if tb.Inside(c.rect) {
			return c.area, nil
		}
	}

	return 0, errors.New("no solution found")
}
