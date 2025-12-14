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

func solvePart2(lines []string) int {
	// TODO: Implement part 2
	return 0
}
