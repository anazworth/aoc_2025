package day11

import (
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(11, Solution{})
}

func (s Solution) Part1(input string) string {
	nodes := parse(input)
	sum := 0

	stack := []string{"you"}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current == "out" {
			sum++
		}
		nextSlice := nodes[current]
		for _, next := range nextSlice {
			stack = append(stack, next)
		}
	}

	return strconv.Itoa(sum)
}

type key struct {
	node    string
	seenDac bool
	seenFft bool
}

func (s Solution) Part2(input string) string {
	nodes := parse(input)

	memo := make(map[key]int)

	var count func(k key) int
	count = func(k key) int {
		if v, ok := memo[k]; ok {
			return v
		}

		// reached end
		if k.node == "out" {
			if k.seenDac && k.seenFft {
				memo[k] = 1
				return 1
			}
			memo[k] = 0
			return 0
		}

		sum := 0
		for _, next := range nodes[k.node] {
			nextKey := key{
				node:    next,
				seenDac: k.seenDac || next == "dac",
				seenFft: k.seenFft || next == "fft",
			}
			sum += count(nextKey)
		}

		memo[k] = sum
		return sum
	}

	result := count(key{node: "svr"})
	return strconv.Itoa(result)
}

func parse(input string) map[string][]string {
	lines := utils.Lines(input)
	result := make(map[string][]string, len(lines))

	for _, line := range lines {
		// Split "aaa: you hhh" → ["aaa", " you hhh"]
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue // or handle error
		}

		name := strings.TrimSpace(parts[0])
		childrenPart := strings.TrimSpace(parts[1])

		var children []string
		if childrenPart != "" {
			children = strings.Fields(childrenPart)
		}

		result[name] = children
	}

	return result
}
