package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

const (
	startString  = "S"
	splitterChar = `\^`
)

var reNext = regexp.MustCompile(splitterChar)

func updateCurrent(current []int, lineLength, i int) {
	switch i {
	case 0:
		current[i], current[i+1] = 0, 1
	case lineLength - 1:
		current[i-1], current[i] = 1, 0
	default:
		current[i-1], current[i], current[i+1] = 1, 0, 1
	}
}

func updateWorldCount(current []int, lineLength, i int) {
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
}

type updateFunc func(current []int, lineLength, i int)

func solver(lines []string, updater updateFunc, countInline bool) int {
	// Keep info about two lines. The current and the next.
	lineLength := len(lines[0])
	current := make([]int, lineLength)
	current[strings.Index(lines[0], startString)] = 1
	next := make([]int, lineLength)
	count := 0
	for _, line := range lines[1:] {
		nextIdx := reNext.FindAllStringIndex(line, -1)
		clear(next)
		if len(nextIdx) > 0 {
			for _, splitter := range nextIdx {
				next[splitter[0]] = 1
			}
			for i := range lineLength {
				if current[i] > 0 && next[i] == 1 {
					if countInline {
						count++
					}
					updater(current, lineLength, i)
				}
			}
		}
	}
	if !countInline {
		for _, v := range current {
			count += v
		}
	}
	return count
}
func solvePart1(lines []string) int {
	return solver(lines, updateCurrent, true)
}

func solvePart2(lines []string) int {
	return solver(lines, updateWorldCount, false)
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
