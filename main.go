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
	"sync"
	"time"
)

const (
	colorReset = "\033[0m"
	colorGray  = "\033[90m"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	days := day.AllDays()
	results := make([]string, len(days)*2) // Store results in order

	for i, d := range days {
		wg.Go(func() {
			dayNum := i + 1
			input := utils.ReadInput(fmt.Sprintf("day/day%02d/input.txt", dayNum))

			start := time.Now()
			result1 := d.Part1(input)
			elapsed1 := time.Since(start)
			line1 := fmt.Sprintf("Day %d - Part 1: %s %s(%v)%s\n", dayNum, result1, colorGray, elapsed1, colorReset)

			start = time.Now()
			result2 := d.Part2(input)
			elapsed2 := time.Since(start)
			line2 := fmt.Sprintf("Day %d - Part 2: %s %s(%v)%s\n", dayNum, result2, colorGray, elapsed2, colorReset)

			mu.Lock()
			results[i*2] = line1
			results[i*2+1] = line2
			mu.Unlock()
		})
	}

	wg.Wait()

	// Print results in order
	for _, line := range results {
		fmt.Print(line)
	}
}
