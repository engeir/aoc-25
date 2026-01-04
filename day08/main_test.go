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

	got := solvePart1(lines, 10)
	want := 40
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T)  {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	
	got := solvePart2(lines)
	want := 25272
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
