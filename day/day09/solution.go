package day9

import (
	"sort"
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
	fill := buildFillWithPerimeter(tiles)

	maxArea := 0
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			if rectangleValid(tiles[i], tiles[j], fill) {
				maxArea = max(maxArea, area(tiles[i], tiles[j]))
			}
		}
	}
	return strconv.Itoa(maxArea)
}

func buildFillWithPerimeter(polygon []pos) map[int][]interval {
	fill := buildFill(polygon)
	perim := buildPerimeterIntervals(polygon)

	for y, ps := range perim {
		fill[y] = mergeIntervals(fill[y], ps)
	}

	return fill
}

func mergeIntervals(a, b []interval) []interval {
	all := append(a, b...)
	if len(all) == 0 {
		return nil
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].x1 < all[j].x1
	})

	out := []interval{all[0]}
	for _, in := range all[1:] {
		last := &out[len(out)-1]
		if in.x1 <= last.x2+1 {
			last.x2 = max(last.x2, in.x2)
		} else {
			out = append(out, in)
		}
	}
	return out
}

func rectangleValid(p1, p2 pos, fill map[int][]interval) bool {
	x1, x2 := min(p1.x, p2.x), max(p1.x, p2.x)
	y1, y2 := min(p1.y, p2.y), max(p1.y, p2.y)

	// top & bottom edges
	if !intervalContains(fill[y1], x1, x2) {
		return false
	}
	if !intervalContains(fill[y2], x1, x2) {
		return false
	}

	// vertical edges
	for y := y1; y <= y2; y++ {
		if !intervalContains(fill[y], x1, x1) {
			return false
		}
		if !intervalContains(fill[y], x2, x2) {
			return false
		}
	}

	return true
}

func intervalContains(intervals []interval, x1, x2 int) bool {
	for _, in := range intervals {
		if in.x1 <= x1 && x2 <= in.x2 {
			return true
		}
	}
	return false
}

type interval struct {
	x1, x2 int
}

func buildFill(polygon []pos) map[int][]interval {
	fill := make(map[int][]interval)

	minY, maxY := polygon[0].y, polygon[0].y
	for _, p := range polygon {
		minY = min(minY, p.y)
		maxY = max(maxY, p.y)
	}

	for y := minY; y <= maxY; y++ {
		var xs []int

		for i := 0; i < len(polygon); i++ {
			p1 := polygon[i]
			p2 := polygon[(i+1)%len(polygon)]

			if p1.y == p2.y {
				continue
			}
			if p1.y > p2.y {
				p1, p2 = p2, p1
			}
			if y < p1.y || y >= p2.y {
				continue
			}

			x := p1.x + (y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y)
			xs = append(xs, x)
		}

		sort.Ints(xs)

		for i := 0; i+1 < len(xs); i += 2 {
			fill[y] = append(fill[y], interval{xs[i], xs[i+1]})
		}
	}

	return fill
}
func buildPerimeterIntervals(polygon []pos) map[int][]interval {
	perim := make(map[int][]interval)

	for i := 0; i < len(polygon); i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%len(polygon)]

		// Horizontal edge
		if p1.y == p2.y {
			x1, x2 := min(p1.x, p2.x), max(p1.x, p2.x)
			perim[p1.y] = append(perim[p1.y], interval{x1, x2})
		}
	}

	return perim
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
