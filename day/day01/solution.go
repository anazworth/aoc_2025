package day01

import (
	"container/ring"
	"fmt"
	"regexp"
	"strconv"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(1, Solution{})
}

func (s Solution) Part1(input string) string {
	directions := parseRotations(input)

	const dialSize = 100
	dial := ring.New(dialSize)

	for i := range dialSize {
		dial.Value = i
		dial = dial.Next()
	}
	dial = dial.Move(50)

	foundZero := 0

	for _, dir := range directions {
		dial = dial.Move(dir)
		if dial.Value == 0 {
			foundZero = foundZero + 1
		}
	}
	return fmt.Sprintf("%d", foundZero)
}

func (s Solution) Part2(input string) string {
	directions := parseRotations(input)

	const dialSize = 100
	dial := ring.New(dialSize)

	for i := range dialSize {
		dial.Value = i
		dial = dial.Next()
	}
	dial = dial.Move(50)

	foundZero := 0

	for _, dir := range directions {
		if dir < 0 {
			for range -dir {
				dial = dial.Prev()

				if dial.Value == 0 {
					foundZero = foundZero + 1
				}
			}
		}
		if dir >= 0 {
			for range dir {
				dial = dial.Next()

				if dial.Value == 0 {
					foundZero = foundZero + 1
				}
			}
		}
	}
	return fmt.Sprintf("%d", foundZero)
}

func parseRotations(input string) []int {
	lines := utils.Lines(input)
	nums := make([]int, len(lines))

	m := regexp.MustCompile("L|R")

	for i, line := range lines {
		normalizedLine := m.ReplaceAllStringFunc(line, func(match string) string {
			if match == "L" {
				return "-"
			}
			return "" // for 'r'
		})
		number, err := strconv.Atoi(normalizedLine)
		if err != nil {
			fmt.Errorf("could not convert %v to integer", normalizedLine)
		}

		nums[i] = number
	}

	return nums
}

var _ day.Day = Solution{} // compile-time check that Solution implements Day
