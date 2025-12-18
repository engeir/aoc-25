package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestPart1(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart1(lines)
	want := 13
	if got != want {
		t.Fatalf("Got %d, but wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart2(lines)
	want := 43
	if got != want {
		t.Fatalf("Got %d, but wanted %d", got, want)
	}
}
