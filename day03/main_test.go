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
	want := 357
	if got != want {
		t.Fatalf("Got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.ReadLines("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart2(lines)
	want := 3121910778619
	if got != want {
		t.Fatalf("Got %d, wanted %d", got, want)
	}
}
