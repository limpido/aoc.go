package main

import (
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func largest(arr []int, start, end int) (int, int) {
	res, idx := 0, 0
	for i, n := range arr {
		if i < start {
			continue
		}
		if i >= end {
			break
		}
		if n > res {
			res = n
			idx = i
		}
	}
	return res, idx
}

func largestNDigit(arr []int, n_digit int) int {
	res, start := 0, 0
	for nth := range n_digit {
		end := len(arr) - (n_digit - nth - 1)
		n, i := largest(arr, start, end)
		start = i + 1
		res = res*10 + n
	}
	return res
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	arr := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, char := range line {
			row[j] = int(char - '0')
		}
		arr[i] = row
	}
	return arr
}

func (s *solution) Part1(input string) any {
	var res int
	banks := parseInput(input)

	for _, b := range banks {
		res += largestNDigit(b, 2)
	}

	return res
}

func (s *solution) Part2(input string) any {
	var res int
	banks := parseInput(input)

	for _, b := range banks {
		res += largestNDigit(b, 12)
	}

	return res
}

func main() {}
