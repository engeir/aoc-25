package main

import (
	"fmt"
	"log"

	"github.com/engeir/aoc-25/utils"
)

func main() {
	// Test input
	testLines, err := utils.ReadLines("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Test Input ===")
	fmt.Println("Part 1:", solvePart1(testLines))  // Should result in 1227775554
	fmt.Println("Part 2:", solvePart2(testLines))

	// Full input
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Test Input ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func solvePart1(lines []string) int {
	return 0
}

func solvePart2(lines []string) int {
	return 0
}
