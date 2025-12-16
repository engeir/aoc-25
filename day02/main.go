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
	// lines_no_newline := strings.TrimSuffix(lines, "\n")
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

func checkRepeatPattern(fullNumber int, part string, length int) int {
	// First check if the part can be used to construct the full number
	if length%len(part) != 0 {
		return 0
	}
	// Is the full number equal to the part repeated enough to get the length of the
	// full?
	multiple := length / len(part)
	constructed, err := strconv.Atoi(strings.Repeat(part, multiple))
	if err != nil {
		log.Fatal(err)
	}
	if constructed == fullNumber {
		return fullNumber
	}
	return 0
}

func checkRepetition2(i int) int {
	asString := strconv.Itoa(i)
	length := len(asString)
	if length < 2 {
		return 0
	}
	for idx := 1; idx < length; idx++ {
		value := checkRepeatPattern(i, asString[:idx], length)
		if value != 0 {
			return value
		}
	}
	return 0
}

type Options struct {
	Part int
}

func getIntRange(s string, opts Options) int {
	parts := strings.Split(s, "-")
	p1, err := strconv.Atoi(strings.TrimSuffix(parts[0], "\n"))
	if err != nil {
		log.Fatal(err)
	}
	p2, err := strconv.Atoi(strings.TrimSuffix(parts[1], "\n"))
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for i := p1; i < p2+1; i++ {
		switch opts.Part {
		case 1:
			count += checkRepetition(i)
		case 2:
			count += checkRepetition2(i)
		}
	}
	return count
}

func solvePart1(lines string) int {
	sections := strings.Split(lines, ",")
	count := 0
	for _, s := range sections {
		count += getIntRange(s, Options{Part: 1})
	}
	return count
}

func solvePart2(lines string) int {
	sections := strings.Split(lines, ",")
	count := 0
	for _, s := range sections {
		count += getIntRange(s, Options{Part: 2})
	}
	return count
}
