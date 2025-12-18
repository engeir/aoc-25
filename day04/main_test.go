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

func Test_solvePart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.lines); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
