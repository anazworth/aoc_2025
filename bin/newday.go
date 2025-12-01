package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

var tmpl = `package day{{.Num}}

import "github.com/anazworth/aoc_2025/day"

type Solution struct{}

func init() {
	day.Register({{.Num}}, Solution{})
}

func (s Solution) Part1(input string) string {
	return "not implemented"
}

func (s Solution) Part2(input string) string {
	return "not implemented"
}
`

var testTmpl = `package day{{.Num}}

import "testing"

const exampleInput = 

func TestPart1(t *testing.T) {
	s := Solution{}
	expected := ""
	result := s.Part1(exampleInput)

	if result != expected {
		t.Errorf("Part 1() = %v, expected %v", result, expected)
	}
}

func TestPart2(t *testing.T) {
	s := Solution{}
	expected := ""
	result := s.Part2(exampleInput)

	if result != expected {
		t.Errorf("Part 2() = %v, expected %v", result, expected)
	}
}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run newday.go 2")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	dir := fmt.Sprintf("day/day%02d", num)
	os.MkdirAll(dir, 0755)

	// Main solution file
	writeTemplate(filepath.Join(dir, "solution.go"), tmpl, num)

	// Solution test file
	writeTemplate(filepath.Join(dir, "solution_test.go"), testTmpl, num)

	fmt.Println("Created day", num)
}

func writeTemplate(path string, tpl string, num int) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.New("day").Parse(tpl))
	t.Execute(f, struct{ Num int }{Num: num})
}
