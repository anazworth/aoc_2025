package main

import (
	"fmt"

	"github.com/anazworth/aoc_2025/day"
	_ "github.com/anazworth/aoc_2025/day/day01"
	_ "github.com/anazworth/aoc_2025/day/day02"
	_ "github.com/anazworth/aoc_2025/day/day03"
	_ "github.com/anazworth/aoc_2025/day/day04"
	_ "github.com/anazworth/aoc_2025/day/day05"
	"github.com/anazworth/aoc_2025/utils"
)

func main() {
	for i, d := range day.AllDays() {
		if i < 3 {
			continue
		}
		input := utils.ReadInput(fmt.Sprintf("day/day%02d/input.txt", i+1))
		fmt.Printf("Day %d - Part 1: %s\n", i+1, d.Part1(input))
		fmt.Printf("Day %d - Part 2: %s\n", i+1, d.Part2(input))
	}
}
