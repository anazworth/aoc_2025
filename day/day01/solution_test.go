package day01

import "testing"

const exampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

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
	expected := "6"
	result := s.Part2(exampleInput)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}
