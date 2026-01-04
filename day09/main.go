package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

type Coord struct {
	X, Y int
}

func linesToCoord(lines []string) []Coord {
	var coords []Coord
	for _, line := range lines {
		lineParts := strings.Split(line, ",")
		linePartsX, err := strconv.Atoi(lineParts[0])
		if err != nil {
			log.Fatal(err)
		}
		linePartsY, err := strconv.Atoi(lineParts[1])
		if err != nil {
			log.Fatal(err)
		}
		coords = append(coords, Coord{linePartsX, linePartsY})
	}
	return coords
}

type AreaPair struct {
	tile1, tile2 Coord
	Area         int
}

func calculateArea(coord1, coord2 Coord) int {
	var (
		x, y int
	)
	if coord1.X < coord2.X {
		x = coord2.X - coord1.X + 1
	} else {
		x = coord1.X - coord2.X + 1
	}
	if coord1.Y < coord2.Y {
		y = coord2.Y - coord1.Y + 1
	} else {
		y = coord1.Y - coord2.Y + 1
	}
	return x * y
}

func findPairAreas(coords []Coord) []AreaPair {
	var allPairs []AreaPair
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			area := calculateArea(coords[i], coords[j])
			allPairs = append(allPairs, AreaPair{coords[i], coords[j], area})
		}
	}
	return allPairs
}

func insideFence(coord1, coord2 Coord, coords []Coord) bool {
	var (
		tl, tr, bl, br Coord
	)
	if coord1.X < coord2.X {
		tl.X = coord1.X
		bl.X = coord1.X
		tr.X = coord2.X
		br.X = coord2.X
	} else {
		tl.X = coord2.X
		bl.X = coord2.X
		tr.X = coord1.X
		br.X = coord1.X
	}
	if coord1.Y < coord2.Y {
		tl.Y = coord1.Y
		bl.Y = coord2.Y
		tr.Y = coord1.Y
		br.Y = coord2.Y
	} else {
		tl.Y = coord2.Y
		bl.Y = coord1.Y
		tr.Y = coord2.Y
		br.Y = coord1.Y
	}
	for _, corner := range []Coord{tl, tr, bl, br} {
		if !isInsideOrOnFence(corner, coords) {
			return false
		}
	}

	// Check points along all 4 edges
	// Sample points along edges to ensure entire rectangle is inside

	// Top edge (from tl to tr)
	for x := tl.X; x <= tr.X; x++ {
		if !isInsideOrOnFence(Coord{x, tl.Y}, coords) {
			return false
		}
	}

	// Bottom edge (from bl to br)
	for x := bl.X; x <= br.X; x++ {
		if !isInsideOrOnFence(Coord{x, bl.Y}, coords) {
			return false
		}
	}

	// Left edge (from tl to bl)
	for y := tl.Y; y <= bl.Y; y++ {
		if !isInsideOrOnFence(Coord{tl.X, y}, coords) {
			return false
		}
	}

	// Right edge (from tr to br)
	for y := tr.Y; y <= br.Y; y++ {
		if !isInsideOrOnFence(Coord{tr.X, y}, coords) {
			return false
		}
	}

	return true
}

func findRestrictedPairAreas(coords []Coord) []AreaPair {
	var allPairs []AreaPair
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			if insideFence(coords[i], coords[j], coords) {
				area := calculateArea(coords[i], coords[j])
				allPairs = append(allPairs, AreaPair{coords[i], coords[j], area})
			}
		}
	}
	return allPairs
}

func isInsideFence(point Coord, coords []Coord) bool {
	// Point-in-polygon using winding number
	winding := 0

	n := len(coords)
	for i := range n {
		p1 := coords[i]
		p2 := coords[(i+1)%n]

		if p1.Y <= point.Y {
			if p2.Y > point.Y {
				if isLeft(p1, p2, point) > 0 {
					winding++
				}
			}
		} else {
			if p2.Y <= point.Y {
				if isLeft(p1, p2, point) < 0 {
					winding--
				}
			}
		}
	}

	return winding != 0
}

func isLeft(p0, p1, p2 Coord) int {
	return (p1.X-p0.X)*(p2.Y-p0.Y) - (p2.X-p0.X)*(p1.Y-p0.Y)
}

func isInsideOrOnFence(point Coord, coords []Coord) bool {
	// First check if point is on any fence segment
	n := len(coords)
	for i := range n {
		p1 := coords[i]
		p2 := coords[(i+1)%n]

		// Check if point is on the line segment from p1 to p2
		if isOnSegment(p1, p2, point) {
			return true
		}
	}

	// If not on fence, check if inside using winding number
	return isInsideFence(point, coords)
}

func isOnSegment(p1, p2, point Coord) bool {
	// Check if point is on the line segment from p1 to p2

	// First check if point is collinear with p1 and p2
	if isLeft(p1, p2, point) != 0 {
		return false // Not on the line
	}

	// Check if point is within the bounding box of the segment
	minX, maxX := p1.X, p2.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := p1.Y, p2.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	return point.X >= minX && point.X <= maxX && point.Y >= minY && point.Y <= maxY
}

func solvePart1(lines []string) int {
	coords := linesToCoord(lines)
	allPairs := findPairAreas(coords)
	slices.SortFunc(allPairs, func(a, b AreaPair) int { return cmp.Compare(b.Area, a.Area) })
	return allPairs[0].Area
}

func solvePart2(lines []string) int {
	coords := linesToCoord(lines)
	// fence := createFence(coords)
	// area := fillFence(fence, coords)

	allPairs := findRestrictedPairAreas(coords)
	slices.SortFunc(allPairs, func(a, b AreaPair) int { return cmp.Compare(b.Area, a.Area) })
	return allPairs[0].Area
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 9 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
