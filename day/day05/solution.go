package day5

import (
	"sort"
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(5, Solution{})
}

type idRange struct {
	start int
	end   int
}

func (s Solution) Part1(input string) string {
	idRanges, ids := parse(input)
	idRanges = reduceRanges(idRanges)

	sum := 0

	for _, idRange := range idRanges {
		for _, id := range ids {
			if id >= idRange.start && id <= idRange.end {
				sum += 1
			}
		}
	}
	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	idRanges, _ := parse(input)
	idRanges = reduceRanges(idRanges)

	sum := 0

	for _, idRange := range idRanges {
		sum += idRange.end + 1 - idRange.start
	}
	return strconv.Itoa(sum)
}

func reduceRanges(input []idRange) []idRange {
	sort.Slice(input, func(i, j int) bool {
		return input[i].start < input[j].start
	})

	result := make([]idRange, 0)

	currStart := input[0].start
	currEnd := input[0].end

	for _, idR := range input {
		if idR.start <= currEnd {
			currEnd = max(currEnd, idR.end)
			continue
		}
		result = append(result, idRange{start: currStart, end: currEnd})
		currStart = idR.start
		currEnd = idR.end
	}

	result = append(result, idRange{start: currStart, end: currEnd})

	return result
}

func parse(input string) ([]idRange, []int) {
	splitInput := strings.Split(input, "\n\n")
	idRangesStr := utils.Lines(splitInput[0])
	ingredientIdsStr := utils.Lines(splitInput[1])

	idRanges := make([]idRange, len(idRangesStr))

	for i, idRangeStr := range idRangesStr {
		split := strings.Split(idRangeStr, "-")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		idRanges[i] = idRange{start: left, end: right}
	}

	ingredientIds := make([]int, len(ingredientIdsStr))

	for i, num := range ingredientIdsStr {
		ingredientIds[i], _ = strconv.Atoi(num)
	}
	return idRanges, ingredientIds
}
