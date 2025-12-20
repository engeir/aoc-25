package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestSolvePart1(t *testing.T)  {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart1(lines)
	want := 3
	if got != want {
		t.Fatalf("Got %d, wanted %d", got, want)
	}
}

func TestSolvePart2(t *testing.T)  {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart2(lines)
	want := 14
	if got != want {
		t.Fatalf("Got %d, wanted %d", got, want)
	}
}
