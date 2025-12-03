package main

import (
	"fmt"
	"github.com/engeir/aoc-25/utils"
	"log"
)

func main() {
	lines, err := utils.ReadLines("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func solvePart1(lines []string) int {
	// TODO: Implement part 1
	return 0
}

func solvePart2(lines []string) int {
	// TODO: Implement part 2
	return 0
}
