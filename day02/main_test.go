package main

import (
	"testing"

	"github.com/engeir/aoc-25/utils"
)

func TestPart1(t *testing.T) {
	lines, err := utils.ReadInput("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart1(lines)
	want := 1227775554
	if got != want {
		t.Errorf("We got %d, wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines, err := utils.ReadInput("./test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := solvePart2(lines)
	want := 4174379265
	if got != want {
		t.Errorf("We got %d, wanted %d", got, want)
	}
}
