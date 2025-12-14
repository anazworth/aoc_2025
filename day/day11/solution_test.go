package day11

import "testing"

const exampleInput = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

const exampleInput2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func TestPart1(t *testing.T) {
	s := Solution{}
	expected := "5"
	result := s.Part1(exampleInput)

	if result != expected {
		t.Errorf("Part 1() = %v, expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := Solution{}
	expected := "2"
	result := s.Part2(exampleInput2)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}
