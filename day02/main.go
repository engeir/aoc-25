package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func main() {
	lines, err := utils.ReadInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Full Input ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}

func checkRepetition(i int) int {
	asString := strconv.Itoa(i)
	length := len(asString)
	if length%2 != 0 {
		return 0
	}
	first := asString[:length/2]
	second := asString[length/2:]
	if first == second {
		return i
	}
	return 0
}

func getIntRange(s string) int {
	parts := strings.Split(s, "-")
	p1, _ := strconv.Atoi(parts[0])
	p2, _ := strconv.Atoi(parts[1])
	count := 0
	for i := p1; i < p2+1; i++ {
		count += checkRepetition(i)
	}
	return count
}

func solvePart1(lines string) int {
	sections := strings.Split(lines, ",")
	count := 0
	for _, s := range sections {
		count += getIntRange(s)
	}
	return count
}

func solvePart2(lines string) int {
	return 0
}
