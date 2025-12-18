package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func checkTopRight(i, j int, lines []string) bool {
	fmt.Println("Here we are")
	nearby_tmp := []string{string(lines[i][j-1]), string(lines[i+1][j-1]), string(lines[i+1][j])}
	nearby := strings.Join(nearby_tmp, "")
	fmt.Println(strings.Count(nearby, "@"))
	return true
}

func checkTopLeft(i, j int, lines []string) bool {
	nearby_tmp := []string{string(lines[i][j+1]), string(lines[i+1][j+1]), string(lines[i+1][j])}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkBottomRight(i, j int, lines []string) bool {
	nearby_tmp := []string{string(lines[i-1][j]), string(lines[i-1][j-1]), string(lines[i][j-1])}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkBottomLeft(i, j int, lines []string) bool {
	nearby_tmp := []string{string(lines[i-1][j]), string(lines[i-1][j+1]), string(lines[i][j+1])}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkTop(i, j int, lines []string) bool {
	nearby_tmp := []string{
		string(lines[i][j-1]),
		string(lines[i+1][j-1]),
		string(lines[i+1][j]),
		string(lines[i+1][j+1]),
		string(lines[i][j+1]),
	}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkRight(i, j int, lines []string) bool {
	nearby_tmp := []string{
		string(lines[i-1][j]),
		string(lines[i-1][j-1]),
		string(lines[i][j-1]),
		string(lines[i+1][j-1]),
		string(lines[i+1][j]),
	}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkBottom(i, j int, lines []string) bool {
	nearby_tmp := []string{
		string(lines[i][j-1]),
		string(lines[i-1][j-1]),
		string(lines[i-1][j]),
		string(lines[i-1][j+1]),
		string(lines[i][j+1]),
	}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkLeft(i, j int, lines []string) bool {
	nearby_tmp := []string{
		string(lines[i-1][j]),
		string(lines[i-1][j+1]),
		string(lines[i][j+1]),
		string(lines[i+1][j+1]),
		string(lines[i+1][j]),
	}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func checkMiddle(i, j int, lines []string) bool {
	nearby_tmp := []string{
		string(lines[i+1][j+1]),
		string(lines[i+1][j-1]),
		string(lines[i+1][j]),
		string(lines[i-1][j+1]),
		string(lines[i-1][j-1]),
		string(lines[i-1][j]),
		string(lines[i][j+1]),
		string(lines[i][j-1]),
	}
	nearby := strings.Join(nearby_tmp, "")
	return strings.Count(nearby, "@") < 4
}

func findGridFamily(i, j, h, w int, lines []string) bool {
	switch {
	// Top right corner
	case i == 0 && j == w-1:
		return checkTopRight(i, j, lines)
	// Top left corner
	case i == 0 && j == 0:
		return checkTopLeft(i, j, lines)
	// Bottom right corner
	case i == h-1 && j == w-1:
		return checkBottomRight(i, j, lines)
	// Bottom left corner
	case i == h-1 && j == 0:
		return checkBottomLeft(i, j, lines)
	// Top
	case i == 0:
		return checkTop(i, j, lines)
	// Right
	case j == w-1:
		return checkRight(i, j, lines)
	// Bottom
	case i == h-1:
		return checkBottom(i, j, lines)
	// Left
	case j == 0:
		return checkLeft(i, j, lines)
	default:
		return checkMiddle(i, j, lines)
	}
}

func solvePart1(lines []string) int {
	// Get the number of lines (height) and row length (width)
	height := len(lines)
	width := len(lines[0])
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if string(lines[i][j]) == "@" {
				if findGridFamily(i, j, height, width, lines) {
					count++
				}
			}
		}
	}
	return count
}

func solvePart2(lines []string) int {
	return 0
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Part 1: ===")
	fmt.Println(solvePart1(lines))
	fmt.Println("=== Part 2: ===")
	solvePart2(lines)
}
