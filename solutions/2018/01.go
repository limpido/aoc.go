package main

import (
	"strings"
	"strconv"
	
	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	var nums []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}
	return nums
}

func (s *Solver) Part1(input string) any {
	nums := parseInput(input)
	return shared.Sum(nums)
}

func (s *Solver) Part2(input string) any {
	nums := parseInput(input)
	m := make(map[int]bool)
	m[0] = true
	cur := 0
	for {
		for _, n := range nums {
			cur += n
			if m[cur] {
				return cur
			}
			m[cur] = true
		}
	}

	return 0
}

func main() {}
