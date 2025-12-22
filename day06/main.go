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

func getGridSize(lines []string) (int, int) {

	height := len(lines)
	length := 0
	for _, l := range lines {
		if len(l) > length {
			length = len(l)
		}
	}
	return height, length
}

func findNumbers(lines []string, calc []int, idx int) []int {
	numb := ""
	for _, line := range lines {
		if idx < len(line) {
			numb += string(line[idx])
		}
	}
	numb = strings.TrimSpace(numb)
	if numb != "" {
		calcTmp, err := strconv.Atoi(strings.TrimSpace(numb))
		if err != nil {
			log.Fatal(err)
		}
		calc = append(calc, calcTmp)
	}
	return calc
}

func findOperator(lines string, idx int) string {
	op := ""
	if idx >= len(lines) {
		op = " "
	} else {
		op = string(lines[idx])
	}
	return strings.TrimSpace(op)
}

func evaluateTask(calc []int, operation string) int {
	subCount := 0
	for _, calc := range calc {
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
	return subCount
}

func solvePart2(lines []string) int {
	height, length := getGridSize(lines)
	count := 0
	calc := make([]int, 0)
	for i := length; i >= 0; i-- {
		calc = findNumbers(lines[:height-1], calc, i)
		op := findOperator(lines[height-1], i)
		if op != "" {
			count += evaluateTask(calc, op)
			calc = make([]int, 0)
		}
	}
	return count
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
