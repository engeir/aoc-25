package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestPart1(t *testing.T) {
	lines, err := utils.ReadLines("./test_input1.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := solvePart1(lines)
	want := 3

	if got != want {
		t.Errorf("solvePart1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.ReadLines("./test_input1.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := solvePart2(lines)
	want := 6

	if got != want {
		t.Errorf("solvePart2() = %d, want %d", got, want)
	}
}
