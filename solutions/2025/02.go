package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func parseInput(input string) [][]int {
	output := [][]int{}
	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		n := shared.Map(strconv.Atoi, strings.Split(r, "-"))
		output = append(output, []int{n[0], n[1]})
	}
	return output
}

func (s *solution) Part1(input string) any {
	var res int
	ranges := parseInput(input)
	for _, rg := range ranges {
		l, r := rg[0], rg[1]
		for n := l; n <= r; n++ {
			s := strconv.Itoa(n)
			size := len(s)
			if size%2 != 0 {
				continue
			}
			if s[:size/2] == s[size/2:] {
				res += n
			}
		}
	}
	return res
}

func isRepeated(s string) bool {
	ss := fmt.Sprintf("%s%s", s, s)
	return strings.Contains(ss[1:len(ss)-1], s)
}

func (s *solution) Part2(input string) any {
	var res int
	ranges := parseInput(input)
	for _, rg := range ranges {
		l, r := rg[0], rg[1]
		for n := l; n <= r; n++ {
			s := strconv.Itoa(n)
			if isRepeated(s) {
				res += n
			}
		}
	}

	return res
}

func main() {}
