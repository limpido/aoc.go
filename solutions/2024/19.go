package main

import (
	"slices"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func parseInput(input string) ([]string, []string) {
	D := strings.Split(input, "\n\n")
	patterns := []string{}
	designs := []string{}
	for _, s := range strings.Split(D[0], ", ") {
		patterns = append(patterns, s)
	}
	for _, d := range strings.Split(D[1], "\n") {
		designs = append(designs, d)
	}
	return patterns, designs
}

func dfs(target string, patterns []string, cache map[string]int) int {
	if v, ok := cache[target]; ok {
		return v
	}
	if len(target) == 0 {
		return 1
	}
	ans := 0
	for _, p := range patterns {
		l := len(p)
		if l > len(target) {
			break
		}
		if target[:l] == p {
			ans += dfs(target[l:], patterns, cache)
		}
	}
	cache[target] = ans
	return ans
}

func (s *solution) Part1(input string) any {
	var res int
	cache := make(map[string]int)
	patterns, designs := parseInput(input)
	slices.SortFunc(patterns, func(s1, s2 string) int {
		return len(s1) - len(s2)
	})
	for _, d := range designs {
		if dfs(d, patterns, cache) > 0 {
			res++
		}
	}
	return res
}

func (s *solution) Part2(input string) any {
	var res int
	cache := make(map[string]int)
	patterns, designs := parseInput(input)
	slices.SortFunc(patterns, func(s1, s2 string) int {
		return len(s1) - len(s2)
	})
	for _, d := range designs {
		res += dfs(d, patterns, cache)
	}
	return res
}

func main() {}
