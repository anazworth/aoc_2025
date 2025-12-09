package day8

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/anazworth/aoc_2025/day"
	"github.com/anazworth/aoc_2025/utils"
)

type Solution struct{}

func init() {
	day.Register(8, Solution{})
}

type jb struct {
	x int
	y int
	z int
}

type connection struct {
	jb1      jb
	jb2      jb
	distance float64
}

func (s Solution) Part1(input string) string {
	jbs := parse(input)
	amtOfPairs := 1000
	if len(jbs) < 50 {
		amtOfPairs = 10
	}

	conns := make([]connection, 0)
	for i := range jbs {
		for j := i + 1; j < len(jbs); j++ {
			dist := distance(jbs[i], jbs[j])
			conns = append(conns, connection{jb1: jbs[i], jb2: jbs[j], distance: dist})
		}
	}
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].distance < conns[j].distance
	})
	uf := NewUnionFind()
	for i := 0; i < amtOfPairs && i < len(conns); i++ {
		uf.Union(conns[i].jb1, conns[i].jb2)
	}

	sizes := uf.GetComponentSizes(jbs)

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	result := sizes[0] * sizes[1] * sizes[2]
	return strconv.Itoa(result)
}

func (s Solution) Part2(input string) string {
	jbs := parse(input)

	conns := make([]connection, 0)
	for i := range jbs {
		for j := i + 1; j < len(jbs); j++ {
			dist := distance(jbs[i], jbs[j])
			conns = append(conns, connection{jb1: jbs[i], jb2: jbs[j], distance: dist})
		}
	}

	sort.Slice(conns, func(i, j int) bool {
		return conns[i].distance < conns[j].distance
	})

	uf := NewUnionFindWithPoints(jbs) // Initialize with all points

	var lastConnection connection
	for _, conn := range conns {
		if uf.Union(conn.jb1, conn.jb2) {
			lastConnection = conn
			if uf.numSets == 1 {
				break
			}
		}
	}

	result := lastConnection.jb1.x * lastConnection.jb2.x
	return strconv.Itoa(result)
}

type UnionFind struct {
	parent  map[jb]jb
	rank    map[jb]int
	numSets int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent:  make(map[jb]jb),
		rank:    make(map[jb]int),
		numSets: 0,
	}
}

func NewUnionFindWithPoints(points []jb) *UnionFind {
	uf := &UnionFind{
		parent:  make(map[jb]jb),
		rank:    make(map[jb]int),
		numSets: len(points),
	}
	for _, p := range points {
		uf.parent[p] = p
		uf.rank[p] = 0
	}
	return uf
}

func (uf *UnionFind) Find(x jb) jb {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
		uf.numSets++
	}
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y jb) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.numSets--
	return true
}

func (uf *UnionFind) GetComponentSizes(allPoints []jb) []int {
	componentMap := make(map[jb]int)

	for _, point := range allPoints {
		root := uf.Find(point)
		componentMap[root]++
	}

	sizes := make([]int, 0, len(componentMap))
	for _, size := range componentMap {
		sizes = append(sizes, size)
	}

	return sizes
}

func distance(a, b jb) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return round(math.Sqrt(dx*dx+dy*dy+dz*dz), 2)
}

func round(x float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(x*multiplier) / multiplier
}

func parse(input string) []jb {
	inputSplit := utils.Lines(input)
	result := make([]jb, len(inputSplit))

	for i, line := range inputSplit {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		result[i] = jb{x: x, y: y, z: z}
	}
	return result
}
