package main

import (
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func (s *solution) Part1(input string) any {
	var res int
	lines := parseInput(input)
	cur := 50
	for _, line := range lines {
		offset, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			cur -= offset
		} else {
			cur += offset
		}
		cur %= 100
		if cur == 0 {
			res++
		}
	}

	return res
}

func (s *solution) Part2(input string) any {
	var res int
	lines := parseInput(input)
	cur := 50
	for _, line := range lines {
		offset, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			for range offset {
				cur++
				cur %= 100
				if cur == 0 {
					res++
				}
			}
		} else {
			for range offset {
				cur--
				cur %= 100
				if cur == 0 {
					res++
				}
			}
		}
	}

	return res
}

func main() {}

