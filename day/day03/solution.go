package day3

import (
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(3, Solution{})
}

func (s Solution) Part1(input string) string {
	banks := parse(input)

	sum := 0

	for _, batts := range banks {
		max := 0
		maxIndex := 0

		for i, num := range batts {
			if i == len(batts)-1 {
				continue
			}
			if num > max {
				max = num
				maxIndex = i
			}
		}

		otherMax := -1

		for i, num := range batts {
			if i <= maxIndex {
				continue
			}
			if num > otherMax {
				otherMax = num
			}
		}

		sum += utils.ConcatInts(max, otherMax)
	}

	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	banks := parse(input)
	sum := 0

	for _, batts := range banks {
		num := findJoltage(batts)
		number := utils.ConcatIntSlice(num)
		sum += number
	}

	return strconv.Itoa(sum)
}

func parse(input string) [][]int {
	lines := utils.Lines(input)
	batteries := make([][]int, 0, len(lines))

	for _, line := range lines {
		nums := make([]int, 0)
		for numStr := range strings.SplitSeq(line, "") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		batteries = append(batteries, nums)
	}
	return batteries
}

func findJoltage(numbers []int) []int {
	result := make([]int, 12)
	left := 0
	for i := range result {
		right := len(numbers) - 12 + i
		maxIndex := utils.FindMaxIndex(numbers[left : right+1])
		actualIndex := left + maxIndex
		result[i] = numbers[actualIndex]
		left = actualIndex + 1
	}
	return result
}
