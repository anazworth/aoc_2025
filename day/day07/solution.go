package day7

import (
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(7, Solution{})
}

func (s Solution) Part1(input string) string {
	manifold := parse(input)
	startPos := strings.Index(manifold[0], "S")
	currentIndices := []int{startPos}

	found := 0

	for _, l := range manifold[1:] {
		current := make(map[int]bool)
		line := strings.Split(l, "")

		for _, idx := range currentIndices {
			switch line[idx] {
			case ".":
				current[idx] = true
			case "^":
				found += 1
				current[idx-1] = true
				current[idx+1] = true
			}
		}
		nextIndices := []int{}
		for k := range current {
			if k >= 0 && k < len(line) {
				nextIndices = append(nextIndices, k)
			}
		}
		currentIndices = nextIndices
	}
	return strconv.Itoa(found)
}

func (s Solution) Part2(input string) string {
	type pos struct {
		index int
		value int
	}
	manifold := parse(input)
	startPos := strings.Index(manifold[0], "S")
	currentIndices := []pos{{index: startPos, value: 1}}

	for _, l := range manifold[1:] {
		current := make(map[int]int)
		line := strings.Split(l, "")

		for _, idx := range currentIndices {
			switch line[idx.index] {
			case ".":
				current[idx.index] += idx.value
			case "^":
				current[idx.index-1] += idx.value
				current[idx.index+1] += idx.value
			}
		}
		nextIndices := []pos{}
		for k, v := range current {
			if k >= 0 && k < len(line) {
				nextIndices = append(nextIndices, pos{index: k, value: v})
			}
		}
		currentIndices = nextIndices
	}
	sum := 0
	for _, p := range currentIndices {
		sum += p.value
	}
	return strconv.Itoa(sum)
}

func parse(input string) []string {
	return utils.Lines(input)
}
