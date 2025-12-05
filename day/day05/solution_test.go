package day5

import "testing"

const exampleInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	s := Solution{}
	expected := "3"
	result := s.Part1(exampleInput)

	if result != expected {
		t.Errorf("Part 1() = %v, expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := Solution{}
	expected := "14"
	result := s.Part2(exampleInput)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}
