package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/engeir/aoc-25/utils"
)

func main() {
	lines, err := utils.ReadLines("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 1 ===")
	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func rotate(line string, pos int) int {
	dir := string(line[0])
	steps, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Printf("Error parsing as a number: %v", err)
		return 0
	}
	switch dir {
	case "L":
		pos -= steps
	case "R":
		pos += steps
	}
	pos += 100
	return pos % 100
}

func solvePart1(lines []string) int {
	start := 50
	count := 0
	for _, line := range lines {
		start = rotate(line, start)
		if start == 0 {
			count++
		}
	}
	return count
}

func rotate_v2(dir string, steps int, pos int) int {
	switch dir {
	case "L":
		pos -= steps
	case "R":
		pos += steps
	}
	return (pos%100 + 100) % 100
}

func floorDiv(a, b int) int {
	if (a < 0) != (b < 0) && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func solvePart2(lines []string) int {
	start := 50
	count := 0
	for _, line := range lines {
		dir := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing as a number: %v", err)
			return 0
		}
		switch dir {
		case "L":
			// Count how many times we land on 0 going left
			// We land on 0 when (start - k) is a multiple of 100, for k in [1, steps]
			// Number of multiples of 100 in range [start - steps, start - 1]
			count += floorDiv(start-1, 100) - floorDiv(start-steps-1, 100)
		case "R":
			// Count how many times we land on 0 going right
			// We land on 0 when (start + k) is a multiple of 100, for k in [1, steps]
			// Number of multiples of 100 in range [start + 1, start + steps]
			count += (start+steps)/100 - start/100
		}
		start = rotate_v2(dir, steps, start)
	}
	return count
}
