package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func findSplitters(line string) [][]int {
	reNext := regexp.MustCompile(`\^`)
	return reNext.FindAllStringIndex(line, -1)
}

func updateCurrent(current []int, lineLength, i int) []int {
	switch i {
	case 0:
		current[i], current[i+1] = 0, 1
	case lineLength - 1:
		current[i-1], current[i] = 1, 0
	default:
		current[i-1], current[i], current[i+1] = 1, 0, 1
	}
	return current
}

func solvePart1(lines []string) int {
	// Keep info about two lines. The current and the next.
	lineLength := len(lines[0])
	current := make([]int, lineLength)
	current[strings.Index(lines[0], "S")] = 1
	count := 0
	for _, line := range lines[1:] {
		nextIdx := findSplitters(line)
		next := make([]int, lineLength)
		if len(nextIdx) != 0 {
			for _, splitter := range nextIdx {
				next[splitter[0]] = 1
			}
			for i := range lineLength {
				if current[i] == next[i] && next[i] == 1 {
					count++
					current = updateCurrent(current, lineLength, i)
				}
			}
		}
	}
	return count
}

func updateWorldCount(current []int, lineLength, i int) []int {
	switch i {
	case 0:
		current[i+1] += current[i]
		current[i] = 0
	case lineLength - 1:
		current[i-1] += current[i]
		current[i] = 0
	default:
		current[i+1] += current[i]
		current[i-1] += current[i]
		current[i] = 0
	}
	return current
}

func solvePart2(lines []string) int {
	// Keep info about two lines. The current and the next.
	lineLength := len(lines[0])
	current := make([]int, lineLength)
	current[strings.Index(lines[0], "S")] = 1
	for _, line := range lines[1:] {
		nextIdx := findSplitters(line)
		next := make([]int, lineLength)
		if len(nextIdx) != 0 {
			for _, splitter := range nextIdx {
				next[splitter[0]] = 1
			}
			for i := range lineLength {
				if current[i] > 0 && next[i] == 1 {
					current = updateWorldCount(current, lineLength, i)
				}
			}
		}
	}
	count := 0
	for _, v := range current {
		count += v
	}
	return count
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 7 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
