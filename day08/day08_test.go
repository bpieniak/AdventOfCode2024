package main

import (
	"strings"
	"testing"
)

const (
	example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	part1Solution = 14
	part2Solution = 34

	example2 = `T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`
	example2Solution = 9
)

func TestPart1(t *testing.T) {
	got := Part1(strings.NewReader(example))

	if got != part1Solution {
		t.Errorf("expected: %d, got: %d", part1Solution, got)
	}
}

func TestPart2Example1(t *testing.T) {
	got := Part2(strings.NewReader(example))

	if got != part2Solution {
		t.Errorf("expected: %d, got: %d", part2Solution, got)
	}
}

func TestPart2Example2(t *testing.T) {
	got := Part2(strings.NewReader(example2))

	if got != example2Solution {
		t.Errorf("expected: %d, got: %d", example2Solution, got)
	}
}
