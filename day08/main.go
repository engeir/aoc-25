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
	X, Y, Z int
}

func parseCoord(pos string) (Coord, error) {
	posIdx := strings.Split(pos, ",")
	posX, err := strconv.Atoi(posIdx[0])
	if err != nil {
		return Coord{}, err
	}
	posY, err := strconv.Atoi(posIdx[1])
	if err != nil {
		return Coord{}, err
	}
	posZ, err := strconv.Atoi(posIdx[2])
	if err != nil {
		return Coord{}, err
	}
	return Coord{posX, posY, posZ}, nil
}

func calcSquaredDistance(pos1, pos2 Coord) int {
	posX := int(pos2.X - pos1.X)
	posY := int(pos2.Y - pos1.Y)
	posZ := int(pos2.Z - pos1.Z)
	return posX*posX + posY*posY + posZ*posZ
}

func RemoveIndex(s [][]Coord, index int) [][]Coord {
	return append(s[:index], s[index+1:]...)
}

type BoolFloatTuple struct {
	LocalOther bool
	Distance   int
}

type UnionFind struct {
	parent map[Coord]Coord
	rank   map[Coord]int
}

// Initialize - each coordinate is its own circuit
func NewUnionFind(coords []Coord) *UnionFind {
	uf := &UnionFind{
		parent: make(map[Coord]Coord),
		rank:   make(map[Coord]int),
	}
	for _, c := range coords {
		uf.parent[c] = c // each coordinate is its own parent initially
		uf.rank[c] = 0
	}
	return uf
}

// Find the root (representative) of this coordinate's circuit
func (uf *UnionFind) Find(c Coord) Coord {
	if uf.parent[c] != c {
		uf.parent[c] = uf.Find(uf.parent[c]) // path compression
	}
	return uf.parent[c]
}

// Are these two coordinates in the same circuit?
func (uf *UnionFind) Same(c1, c2 Coord) bool {
	return uf.Find(c1) == uf.Find(c2)
}

// Merge the circuits containing these coordinates
func (uf *UnionFind) Union(c1, c2 Coord) {
	root1 := uf.Find(c1)
	root2 := uf.Find(c2)

	if root1 == root2 {
		return // already in same circuit
	}

	// Union by rank
	if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
	} else if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
	} else {
		uf.parent[root2] = root1
		uf.rank[root1]++
	}
}

type DistPair struct {
	distance int
	coord1   Coord
	coord2   Coord
}

// Parse coordinates
func getCoords(lines []string) []Coord {
	var coords []Coord
	for _, line := range lines {
		parsed, err := parseCoord(line)
		if err != nil {
			log.Fatal(err)
		}
		coords = append(coords, parsed)
	}
	return coords
}

func computePairwiseDistances(coords []Coord) []DistPair {
	var allPairs []DistPair
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			dist := calcSquaredDistance(coords[i], coords[j])
			allPairs = append(allPairs, DistPair{
				distance: dist,
				coord1:   coords[i],
				coord2:   coords[j],
			})
		}
	}
	return allPairs
}

func extractUniqueDistances(allPairs []DistPair) []int {
	uniqueDistances := []int{}
	lastDist := -1
	for _, pair := range allPairs {
		if pair.distance != lastDist {
			uniqueDistances = append(uniqueDistances, pair.distance)
			lastDist = pair.distance
		}
	}
	return uniqueDistances
}

func solvePart1(lines []string, maxConnections int) int {
	coords := getCoords(lines)
	allPairs := computePairwiseDistances(coords)
	slices.SortFunc(allPairs, func(a, b DistPair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	uniqueDistances := extractUniqueDistances(allPairs)
	uf := NewUnionFind(coords)

	// For each nth iteration
	for nth := 1; nth <= maxConnections; nth++ {
		if nth > len(uniqueDistances) {
			break
		}

		nthDistance := uniqueDistances[nth-1] // 0-indexed

		// Check if the nth distance is inter-circuit
		isInterCircuit := false
		for _, pair := range allPairs {
			if pair.distance == nthDistance {
				if !uf.Same(pair.coord1, pair.coord2) {
					isInterCircuit = true
					break
				}
			} else if pair.distance > nthDistance {
				break
			}
		}

		// If inter-circuit, merge the closest inter-circuit pair
		if isInterCircuit {
			for _, pair := range allPairs {
				if !uf.Same(pair.coord1, pair.coord2) {
					uf.Union(pair.coord1, pair.coord2)
					break
				}
			}
		}
	}

	// Count circuit sizes using Union-Find
	circuitSizes := make(map[Coord]int)
	for _, coord := range coords {
		root := uf.Find(coord)
		circuitSizes[root]++
	}

	// Extract sizes and sort
	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	slices.SortFunc(sizes, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	// Multiply the 3 largest unique sizes
	result := 1
	lastLength := 0
	count := 0

	for _, size := range sizes {
		if count >= 3 {
			break
		}
		if lastLength != size {
			count++
			lastLength = size
			result *= size
		}
	}

	return result
}

func solvePart2(lines []string) int {
	coords := getCoords(lines)
	allPairs := computePairwiseDistances(coords)
	slices.SortFunc(allPairs, func(a, b DistPair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	uf := NewUnionFind(coords)
	var lastPair DistPair
	for _, pair := range allPairs {
		if !uf.Same(pair.coord1, pair.coord2) {
			lastPair = pair
			uf.Union(pair.coord1, pair.coord2)
		}
	}

	return lastPair.coord1.X * lastPair.coord2.X
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 8 ===")
	fmt.Println("Part 1:", solvePart1(lines, 1000))
	fmt.Println("Part 2:", solvePart2(lines))
}
