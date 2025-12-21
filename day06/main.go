package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func createMatrix(lines []string) [][]int {
	factors := len(lines)
	tasks := strings.Fields(lines[factors-1])
	matrix := make([][]int, len(tasks))
	for i := range tasks {
		matrix[i] = make([]int, factors)
	}
	for t, line := range lines {
		taskElements := strings.Fields(line)
		for f, factor := range taskElements {
			result, err := strconv.Atoi(string(factor))
			if err != nil {
				log.Fatal(err)
			}
			matrix[f][t] = result
		}
	}
	return matrix
}

func evaluateMatrix(parts [][]int, operations []string) int {
	count := 0
	for i, operation := range operations {
		calculation := parts[i]
		subCount := 0
		for _, calc := range calculation {
			if subCount == 0 {
				subCount = calc
				continue
			}
			switch operation {
			case "*":
				subCount *= calc
			case "+":
				subCount += calc
			}
		}
		count += subCount
	}
	return count
}

func solvePart1(lines []string) int {
	elemments := len(lines) - 1
	operations := strings.Fields(lines[elemments])
	parts := createMatrix(lines[:elemments])
	return evaluateMatrix(parts, operations)
}

func solvePart2(lines []string) int {
	return 0
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 6 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
