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
	want := 7
	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
