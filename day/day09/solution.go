package day9

import (
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(9, Solution{})
}

type pos struct {
	x int
	y int
}

func (s Solution) Part1(input string) string {
	tiles := parse(input)
	max := -1
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			curr := area(tiles[i], tiles[j])
			if curr > max {
				max = curr
			}
		}
	}
	return strconv.Itoa(max)
}
func (s Solution) Part2(input string) string {
	tiles := parse(input)
	floormap := getAllPositions(tiles)

	// Pre-compute all interior points
	interiorMap := computeInteriorPoints(tiles, floormap)

	maxArea := findMaxRectangleArea(tiles, interiorMap)
	return strconv.Itoa(maxArea)
}

func computeInteriorPoints(polygon []pos, perimeter map[pos]bool) map[pos]bool {
	interior := make(map[pos]bool)

	// Find bounding box
	minX, maxX := polygon[0].x, polygon[0].x
	minY, maxY := polygon[0].y, polygon[0].y
	for _, p := range polygon {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// Check every point in bounding box once
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := pos{x, y}
			if perimeter[p] || isInside(p, polygon) {
				interior[p] = true
			}
		}
	}

	return interior
}

func findMaxRectangleArea(tiles []pos, interior map[pos]bool) int {
	maxArea := 0

	// Try all pairs of points from the tiles list
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			p1 := tiles[i]
			p2 := tiles[j]

			// Check if the rectangle is valid using pre-computed interior
			if isRectangleValidFast(p1, p2, interior) {
				curr := area(p1, p2)
				if curr > maxArea {
					maxArea = curr
				}
			}
		}
	}

	return maxArea
}

func isRectangleValidFast(p1, p2 pos, interior map[pos]bool) bool {
	x1, x2 := min(p1.x, p2.x), max(p1.x, p2.x)
	y1, y2 := min(p1.y, p2.y), max(p1.y, p2.y)

	// Check all positions on the rectangle's perimeter
	// Top and bottom edges
	for x := x1; x <= x2; x++ {
		if !interior[pos{x, y1}] || !interior[pos{x, y2}] {
			return false
		}
	}

	// Left and right edges (excluding corners already checked)
	for y := y1 + 1; y < y2; y++ {
		if !interior[pos{x1, y}] || !interior[pos{x2, y}] {
			return false
		}
	}

	return true
}

func isInside(point pos, polygon []pos) bool {
	count := 0
	n := len(polygon)

	for i := 0; i < n; i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%n]

		if rayIntersectsSegment(point, p1, p2) {
			count++
		}
	}

	return count%2 == 1
}

func rayIntersectsSegment(point, p1, p2 pos) bool {
	if p1.y == p2.y {
		return false
	}

	if point.y < min(p1.y, p2.y) || point.y >= max(p1.y, p2.y) {
		return false
	}

	xIntersect := p1.x + (point.y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y)

	return xIntersect >= point.x
}

func parse(input string) []pos {
	lines := utils.Lines(input)
	result := make([]pos, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		result[i] = pos{x, y}
	}
	return result
}

func getAllPositions(positions []pos) map[pos]bool {
	allPos := make(map[pos]bool)

	for i := 0; i < len(positions)-1; i++ {
		start := positions[i]
		end := positions[(i+1)%len(positions)] // Wrap to start

		// Get all positions between start and end (inclusive)
		between := getPositionsBetween(start, end)
		for _, p := range between {
			allPos[p] = true
		}
	}

	return allPos
}

func getPositionsBetween(start, end pos) []pos {
	var result []pos

	// Determine direction
	dx := sign(end.x - start.x)
	dy := sign(end.y - start.y)

	// Walk from start to end
	current := start
	for current != end {
		result = append(result, current)
		current.x += dx
		current.y += dy
	}
	result = append(result, end) // Include the end position

	return result
}

func sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
func area(p1, p2 pos) int {
	width := abs(p2.x-p1.x) + 1
	height := abs(p2.y-p1.y) + 1
	return width * height
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
