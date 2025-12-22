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
	want := 40
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	lines, _ := utils.ReadLines("./test_input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart1(lines)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	lines, _ := utils.ReadLines("./test_input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart2(lines)
	}
}
