package main

import (
	"slices"
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseLights(s string) []int {
	var lights []int
	for i, c := range s[1 : len(s)-1] {
		if c == '#' {
			lights = append(lights, i)
		}
	}
	return lights
}

func parseSchematics(ss []string) [][]int {
	var schematics [][]int
	for _, s := range ss {
		s = s[1 : len(s)-1]
		nums := shared.Map(strconv.Atoi, strings.Split(s, ","))
		schematics = append(schematics, nums)
	}
	return schematics
}

func parseJoltages(s string) []int {
	return shared.Map(strconv.Atoi, strings.Split(s[1:len(s)-1], ","))
}

func parseInput(input string) ([][]int, [][][]int, [][]int) {
	lines := strings.Split(input, "\n")
	lights := [][]int{}
	schematics := [][][]int{}
	joltages := [][]int{}
	for _, l := range lines {
		lst := strings.Split(l, " ")
		lights = append(lights, parseLights(lst[0]))
		schematics = append(schematics, parseSchematics(lst[1:len(lst)-1]))
		joltages = append(joltages, parseJoltages(lst[len(lst)-1]))
	}
	return lights, schematics, joltages
}

func fewestPresses(target []int, m map[int]int, presses [][]int, startIdx int, count int, minCount int) int {
	if minCount > 0 && count > minCount {
		return minCount
	}
	found := true
	// on: odd, off: even
	for k, v := range m {
		if slices.Contains(target, k) && v%2 != 1 {
			found = false
			break
		} else if !slices.Contains(target, k) && v%2 == 1 {
			found = false
			break
		}
	}
	if found {
		if minCount > 0 {
			return min(count, minCount)
		}
		return count
	}

	for i := startIdx; i < len(presses); i++ {
		p := presses[i]
		for _, n := range p {
			m[n]++
		}
		c := fewestPresses(target, m, presses, startIdx+1, count+1, minCount)
		if minCount > 0 {
			minCount = min(minCount, c)
		} else {
			minCount = c
		}
		for _, n := range p {
			m[n]--
		}
	}
	return minCount
}

func (s *Solver) Part1(input string) any {
	lights, schematics, _ := parseInput(input)
	var res int
	for i, schs := range schematics {
		m := make(map[int]int)
		for _, light := range lights[i] {
			m[light] = 0
		}
		count := fewestPresses(lights[i], m, schs, 0, 0, 0)
		res += count
	}
	return res
}

func (s *Solver) Part2(input string) any {
	return 0
}

func main() {}
