package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := solvePart1(lines)
	want := 21
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := solvePart2(lines)
	want := 0
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
