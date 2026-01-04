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

func solvePart1(lines []string) int {
	coords := linesToCoord(lines)
	allPairs := findPairAreas(coords)
	slices.SortFunc(allPairs, func(a, b AreaPair) int { return cmp.Compare(b.Area, a.Area) })
	return allPairs[0].Area
}

func solvePart2(lines []string) int {
	return 0
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
