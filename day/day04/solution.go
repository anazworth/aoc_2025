package day4

import (
	"strconv"
	"strings"
	"sync"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(4, Solution{})
}

type point struct {
	x int
	y int
}

var dirs = [8]point{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func (s Solution) Part1(input string) string {
	grid := parseGrid(input)

	sum := 0
	var sumMu sync.Mutex
	var wg sync.WaitGroup

	for y, row := range grid {
		for x, char := range row {
			if char == "@" {
				wg.Go(func() {
					accessible := checkRoll(point{x: x, y: y}, &grid)
					if accessible {
						sumMu.Lock()
						defer sumMu.Unlock()
						sum += 1
					}
				})
			}
		}
	}
	wg.Wait()

	return strconv.Itoa(sum)
}

func (s Solution) Part2(input string) string {
	grid := parseGrid(input)

	sum := 0

	for {
		done := true
		for y, row := range grid {
			for x, char := range row {
				if char == "@" {
					accessible := checkRoll(point{x: x, y: y}, &grid)
					if accessible {
						sum += 1
						done = false
						grid[y][x] = "."
						break
					}
				}
			}
			if !done {
				break
			}
		}
		if done {
			break
		}
	}

	return strconv.Itoa(sum)
}

func checkRoll(roll point, grid *[][]string) bool {
	adjRolls := 0
	g := *grid
	rows := len(g)
	cols := len(g[0])

	for _, d := range dirs {
		ny := roll.y + d.y
		nx := roll.x + d.x

		if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
			continue
		}

		if g[ny][nx] == "@" {
			adjRolls += 1
		}
	}
	return adjRolls < 4
}

func parseGrid(input string) [][]string {
	lines := utils.Lines(input)
	grid := make([][]string, len(lines))

	for i, line := range lines {
		result := make([]string, len(line))
		for j, char := range strings.Split(line, "") {
			if char == "\n" {
				continue
			}
			result[j] = char
		}
		if len(result) > 0 {
			grid[i] = result
		}
	}
	return grid
}
