package day10

import "testing"

const exampleInput = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func TestPart1(t *testing.T) {
	s := Solution{}
	expected := "7"
	result := s.Part1(exampleInput)

	if result != expected {
		t.Errorf("Part 1() = %v, expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := Solution{}
	expected := ""
	result := s.Part2(exampleInput)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}

func TestToggleDiagram(t *testing.T) {
	expected := []bool{false, false, true}
	result := toggleDiagram([]bool{false, true, false}, []int{1, 2})

	if len(result) != len(expected) {
		t.Errorf("toggleDiagram length = %d, expected %d", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("toggleDiagram = %v, expected %v", result, expected)
			return
		}
	}
}
