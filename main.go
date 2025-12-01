package main

import (
	"fmt"

	"github.com/anazworth/aoc_2025/day"
	_ "github.com/anazworth/aoc_2025/day/day01"
	"github.com/anazworth/aoc_2025/utils"
)

func main() {
	for i, d := range day.AllDays() {
		input := utils.ReadInput(fmt.Sprintf("day/day%02d/input.txt", i))
		fmt.Printf("Day %d - Part 1: %s\n", i, d.Part1(input))
		fmt.Printf("Day %d - Part 2: %s\n", i, d.Part2(input))
	}
}
