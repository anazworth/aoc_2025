package main

import (
	"fmt"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/day/day01"
	"github.com/anazworth/aoc_2025/utils"
)

func main() {
	days := []day.Day{
		day01.Solution{},
	}

	for i, d := range days {
		input := utils.ReadInput(fmt.Sprintf("day/day%02d/input.txt", i+1))
		fmt.Printf("Day %d - Part 1: %s\n", i+1, d.Part1(input))
		fmt.Printf("Day %d - Part 2: %s\n", i+1, d.Part2(input))
	}
}
