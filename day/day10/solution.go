package day10

import (
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(10, Solution{})
}

func (s Solution) Part1(input string) string {
	machines := parse(input)

	sum := 0

	for _, m := range machines {
		// minSteps := bfs(m)
		// if minSteps != -1 {
		// 	sum += minSteps
		// }
		sum += solve(m)
	}

	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	return "not implemented"
}

type machine struct {
	diagram    []bool
	schematics [][]int
	joltage    []int
}

type state struct {
	diagram []bool
	used    []int // indices of schematics used
}

// trying out bitmasking
func solve(m machine) int {
	n := len(m.schematics)
	target := m.diagram
	best := math.MaxInt

	for mask := 0; mask < (1 << n); mask++ {
		steps := bits.OnesCount(uint(mask))
		if steps >= best {
			continue
		}

		diagram := make([]bool, len(target))

		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				for _, pos := range m.schematics[i] {
					diagram[pos] = !diagram[pos]
				}
			}
		}

		if slices.Equal(diagram, target) {
			best = steps
		}
	}

	if best == math.MaxInt {
		return -1
	}
	return best
}
func bfs(m machine) int {
	// Start with a blank diagram (all false)
	start := make([]bool, len(m.diagram))

	if slices.Equal(start, m.diagram) {
		return 0
	}

	queue := []state{{diagram: start, used: []int{}}}
	visited := make(map[string]bool)
	visited[stateKey(start, []int{})] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Try adding each schematic that hasn't been used yet
		for i, schematic := range m.schematics {
			// Check if this schematic has already been used in this path
			if contains(current.used, i) {
				continue
			}

			next := toggleDiagram(current.diagram, schematic)
			newUsed := append(slices.Clone(current.used), i)

			// Check if we reached the target
			if slices.Equal(next, m.diagram) {
				return len(newUsed)
			}

			key := stateKey(next, newUsed)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, state{diagram: next, used: newUsed})
			}
		}
	}

	return -1 // No solution found
}

func contains(slice []int, val int) bool {
	return slices.Contains(slice, val)
}

func stateKey(diagram []bool, used []int) string {
	diagramKey := boolSliceKey(diagram)
	usedCopy := make([]int, len(used))
	copy(usedCopy, used)
	slices.Sort(usedCopy)
	return fmt.Sprintf("%s-%v", diagramKey, usedCopy)
}

func boolSliceKey(b []bool) string {
	result := make([]byte, len(b))
	for i, v := range b {
		if v {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}

func toggleDiagram(diagram []bool, schematic []int) []bool {
	result := slices.Clone(diagram)
	for _, val := range schematic {
		result[val] = !result[val]
	}
	return result
}

func parse(input string) []machine {
	lines := utils.Lines(input)
	result := make([]machine, len(lines))

	for i, line := range lines {
		m, _ := parseLine(line)
		result[i] = *m
	}

	return result
}

func parseLine(line string) (*machine, error) {
	result := &machine{}

	// Extract bracket content [...] and convert to []bool
	bracketRe := regexp.MustCompile(`\[([^\]]+)\]`)
	bracketMatch := bracketRe.FindStringSubmatch(line)
	if len(bracketMatch) > 1 {
		for _, ch := range bracketMatch[1] {
			result.diagram = append(result.diagram, ch == '#')
		}
	}

	// Extract all parentheses groups (...) and convert to [][]int
	parenRe := regexp.MustCompile(`\(([^)]+)\)`)
	parenMatches := parenRe.FindAllStringSubmatch(line, -1)
	for _, match := range parenMatches {
		if len(match) > 1 {
			var nums []int
			for _, numStr := range strings.Split(match[1], ",") {
				var num int
				fmt.Sscanf(strings.TrimSpace(numStr), "%d", &num)
				nums = append(nums, num)
			}
			result.schematics = append(result.schematics, nums)
		}
	}

	// Extract braces content {...} and convert to []int
	braceRe := regexp.MustCompile(`\{([^}]+)\}`)
	braceMatch := braceRe.FindStringSubmatch(line)
	if len(braceMatch) > 1 {
		for _, numStr := range strings.Split(braceMatch[1], ",") {
			var num int
			fmt.Sscanf(strings.TrimSpace(numStr), "%d", &num)
			result.joltage = append(result.joltage, num)
		}
	}

	return result, nil
}
