package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func createFreshRanges(lines []string) (map[int][2]int, int) {
	m := make(map[int][2]int)
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
		if len(m) == 0 {
			tmp := [2]int{start, stop}
			m[0] = tmp
			continue
		}
		for k, v := range m {
			switch {
			case start <= v[0] && stop > v[0] && stop < v[1]:
				tmp := [2]int(m[k])
				tmp[0] = start
				m[k] = tmp
			case start <= v[0] && stop >= v[1]:
				tmp := [2]int(m[k])
				tmp[0] = start
				tmp[1] = stop
				m[k] = tmp
			case start > v[0] && stop >= v[1] && start < v[1]:
				tmp := [2]int(m[k])
				tmp[1] = stop
				m[k] = tmp
			case start > v[0] && stop < v[1]:
			default:
				tmp := [2]int{start, stop}
				m[j] = tmp
			}
		}
	}
	return m, lineNumber
}

func concatenateRanges(ranges map[int][2]int) (map[int][2]int, int) {
	count := 0
	for j, startStop := range ranges {
		for k, v := range ranges {
			if j == k {
				continue
			}
			switch {
			case startStop[0] <= v[0] && startStop[1] > v[0] && startStop[1] < v[1]:
				tmp := [2]int(ranges[k])
				tmp[0] = startStop[0]
				ranges[k] = tmp
				delete(ranges, j)
				count++
			case startStop[0] <= v[0] && startStop[1] >= v[1]:
				tmp := [2]int(ranges[k])
				tmp[0] = startStop[0]
				tmp[1] = startStop[1]
				ranges[k] = tmp
				delete(ranges, j)
				count++
			case startStop[0] > v[0] && startStop[1] >= v[1] && startStop[0] < v[1]:
				tmp := [2]int(ranges[k])
				tmp[1] = startStop[1]
				ranges[k] = tmp
				delete(ranges, j)
				count++
			case startStop[0] > v[0] && startStop[1] < v[1]:
				delete(ranges, j)
				count++
			}
		}
	}
	return ranges, count
}

func countFreshIngredients(lines []string, lineNumber int, fresh map[int][2]int) int {
	count := 0
	for _, line := range lines[lineNumber:] {
		result, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range fresh {
			if result > v[0] && result <= v[1] {
				count++
			}
		}
	}
	return count
}

func solvePart1(lines []string) int {
	m, lineNumber := createFreshRanges(lines)
	c := 0
	for {
		m, c = concatenateRanges(m) 
		if c == 0 {
			break
		}
	}
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
