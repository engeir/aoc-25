package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func findLargestInt(s string) (int, int) {
	idx := 0
	large := 0
	for i, v := range strings.Split(s, "") {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if v_int > large {
			idx = i
			large = v_int
		}
	}
	return idx, large
}

func createNDigitNumber(numb string, size int) int {
	length := len(numb)
	large := ""
	idx := 0
	window := length - size + 1
	for i := 0; i < size; i++ {
		idx_tmp := 0
		large_i := 0
		numb_part := numb[idx:idx+window]
		idx_tmp, large_i = findLargestInt(numb_part)
		idx += 1 + idx_tmp
		window -= idx_tmp
		large_s := strconv.Itoa(large_i)
		large_tmp := []string{large, large_s}
		large = strings.Join(large_tmp, "")
	}
	numb_i, err := strconv.Atoi(large)
	if err != nil {
		log.Fatal(err)
	}
	return numb_i
}

func solvePart1(lines []string) int {
	count := 0
	for _, v := range lines {
		count += createNDigitNumber(v, 2)
	}
	return count
}
func solvePart2(lines []string) int {
	count := 0
	for _, v := range lines {
		count += createNDigitNumber(v, 12)
	}
	return count
}
func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Day 3 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
