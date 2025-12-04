package day3

import "testing"

const exampleInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	s := Solution{}
	expected := "357"
	result := s.Part1(exampleInput)

	if result != expected {
		t.Errorf("Part 1() = %v, expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := Solution{}
	expected := "3121910778619"
	result := s.Part2(exampleInput)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}
