package day6

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(6, Solution{})
}

func (s Solution) Part1(input string) string {
	ws := parse(input)
	ops := ws[len(ws)-1]
	ws = ws[:len(ws)-1]

	sum := 0

	for i, op := range ops {
		localSum, _ := strconv.Atoi(ws[0][i])
		for _, row := range ws[1:] {
			val, _ := strconv.Atoi(row[i])
			switch op {
			case "*":
				localSum *= val
			case "+":
				localSum += val
			}
		}
		sum += localSum
	}

	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	ws := strings.Split(input, "\n")
	if len(ws[len(ws)-1]) == 0 {
		ws = ws[:len(ws)-1]
	}

	sum := 0
	currentCalc := make([]int, 0)

	for col := len(ws[0]) - 1; col >= 0; col-- {
		buffer := ""
		for row := range ws {
			val := string(ws[row][col])
			if val == " " {
				continue
			}
			if val == "*" || val == "+" {
				if len(buffer) > 0 {
					num, _ := strconv.Atoi(buffer)
					currentCalc = append(currentCalc, num)
					buffer = ""
				}
				sum += doOp(currentCalc, val)
				currentCalc = make([]int, 0)
				continue
			}
			buffer += val
		}
		if len(buffer) > 0 {
			num, _ := strconv.Atoi(buffer)
			currentCalc = append(currentCalc, num)
		}
	}

	return strconv.Itoa(sum)
}

func doOp(input []int, op string) int {
	sum := input[0]

	for _, val := range input[1:] {
		switch op {
		case "*":
			sum *= val
		case "+":
			sum += val
		}
	}
	return sum
}

func parse(input string) [][]string {
	lines := utils.Lines(input)
	re := regexp.MustCompile(`[\d*+]+`)
	result := make([][]string, len(lines))

	for i, line := range lines {
		matches := re.FindAllString(line, -1)
		result[i] = matches
	}
	return result
}
