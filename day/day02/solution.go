package day2

import (
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"

	"github.com/dlclark/regexp2"
)

type Solution struct{}

func init() {
	day.Register(2, Solution{})
}

func (s Solution) Part1(input string) string {
	sum := 0

	m := regexp2.MustCompile(`^(.*)\1$`, 0)

	ranges := parse(input)
	for _, idRange := range ranges {
		for id := idRange.start; id <= idRange.end; id++ {
			match, _ := m.MatchString(strconv.Itoa(id))
			if match {
				sum = sum + id
			}
		}
	}
	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	return "not implemented"
}

type idRange struct {
	start int
	end   int
}

func parse(input string) []idRange {
	ranges := []idRange{}

	input = strings.ReplaceAll(input, "\n", "")

	for pair := range strings.SplitSeq(input, ",") {
		pairStr := strings.Split(pair, "-")
		left, err := strconv.Atoi(pairStr[0])
		right, err := strconv.Atoi(pairStr[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, idRange{start: left, end: right})
	}

	return ranges
}
