package utils

import (
	"strconv"
	"strings"
)

func ConcatInts(x, y int) int {
	s := strconv.Itoa(x) + strconv.Itoa(y)
	result, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}
	return result
}

func ConcatIntSlice(numbers []int) int {
	var sb strings.Builder
	for _, n := range numbers {
		sb.WriteString(strconv.Itoa(n))
	}
	result, _ := strconv.Atoi(sb.String())
	return result
}

// returns max index (first occurence)
func FindMaxIndex(numbers []int) int {
	maxIndex := 0
	for i, val := range numbers {
		if i == 0 {
			continue
		}
		if val > numbers[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
