package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func createFreshRanges(lines []string) (map[int]bool, int) {
	m := make(map[int]bool)
	lineNumber := 0
	for j, line := range lines {
		if line == "" {
			lineNumber = j + 1
			break
		}
		ranges := strings.Split(line, "-")
		start, err := strconv.Atoi(ranges[0])
		if err != nil {
			log.Fatal(err)
		}
		stop, err := strconv.Atoi(ranges[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := start; i < stop+1; i++ {
			m[i] = true
		}
	}
	return m, lineNumber
}

func countFreshIngredients(lines []string, lineNumber int, fresh map[int]bool) int {
	count := 0
	for _, line := range lines[lineNumber:] {
		result, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if fresh[result] {
			count++
		}
	}
	return count
}

func solvePart1(lines []string) int {
	m, lineNumber := createFreshRanges(lines)
	return countFreshIngredients(lines, lineNumber, m)
}

func solvePart2(lines []string) int {
	return 0
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Part 1 ===")
	fmt.Println(solvePart1(lines))
	fmt.Println("=== Part 2 ===")
	fmt.Println(solvePart2(lines))
}
