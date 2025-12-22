package main

import (
	"fmt"
	"log"

	"github.com/engeir/aoc-25/utils"
)

const neighboursRoof = 4

var (
	down      = [2]int{1, 0}
	downLeft  = [2]int{1, -1}
	downRight = [2]int{1, 1}
	left      = [2]int{0, -1}
	right     = [2]int{0, 1}
	up        = [2]int{-1, 0}
	upLeft    = [2]int{-1, -1}
	upRight   = [2]int{-1, 1}
)

var neighbourOffset = map[string][][2]int{
	"top-right": {
		down,
		downLeft,
		left,
	},
	"top-left": {
		down,
		downRight,
		right,
	},
	"bottom-right": {
		left,
		up,
		upLeft,
	},
	"bottom-left": {
		right,
		up,
		upRight,
	},
	"top": {
		down,
		downLeft,
		downRight,
		left,
		right,
	},
	"right": {
		down,
		downLeft,
		left,
		up,
		upLeft,
	},
	"bottom": {
		left,
		right,
		up,
		upLeft,
		upRight,
	},
	"left": {
		down,
		downRight,
		right,
		up,
		upRight,
	},
	"middle": {
		down,
		downLeft,
		downRight,
		left,
		right,
		up,
		upLeft,
		upRight,
	},
}

func checkNeighbours(i, j int, lines []string, offsets [][2]int) bool {
	count := 0
	for _, offset := range offsets {
		row := i + offset[0]
		col := j + offset[1]
		if lines[row][col] == '@' {
			count++
		}
	}
	return count < neighboursRoof
}

func getPosition(i, j, h, w int) string {
	switch {
	// Top right corner
	case i == 0 && j == w-1:
		return "top-right"
	// Top left corner
	case i == 0 && j == 0:
		return "top-left"
	// Bottom right corner
	case i == h-1 && j == w-1:
		return "bottom-right"
	// Bottom left corner
	case i == h-1 && j == 0:
		return "bottom-left"
	// Top
	case i == 0:
		return "top"
	// Right
	case j == w-1:
		return "right"
	// Bottom
	case i == h-1:
		return "bottom"
	// Left
	case j == 0:
		return "left"
	default:
		return "middle"
	}
}

func findGridFamily(i, j, h, w int, lines []string) bool {
	offset := neighbourOffset[getPosition(i, j, h, w)]
	return checkNeighbours(i, j, lines, offset)
}

func solvePart1(lines []string) int {
	// Get the number of lines (height) and row length (width)
	height := len(lines)
	width := len(lines[0])
	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lines[i][j] == '@' {
				if findGridFamily(i, j, height, width, lines) {
					count++
				}
			}
		}
	}
	return count
}

func removeRolls(lines []string) ([]string, int) {
	// Get the number of lines (height) and row length (width)
	height := len(lines)
	width := len(lines[0])
	count := 0
	result := make([]string, len(lines))
	copy(result, lines)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if lines[i][j] == '@' {
				if findGridFamily(i, j, height, width, lines) {
					s := []byte(result[i])
					s[j] = '.'
					result[i] = string(s)
					count++
				}
			}
		}
	}
	return result, count
}

func solvePart2(lines []string) int {
	removableRolls := 0
	count := 0
	for {
		lines, removableRolls = removeRolls(lines)
		count += removableRolls
		if removableRolls == 0 {
			break
		}
	}
	return count
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=== Day 4 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
