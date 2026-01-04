package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	got := solvePart1(lines)
	want := 50
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	got := solvePart2(lines)
	want := 24
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
