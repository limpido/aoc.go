package main

import (
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func count(input string) int {
	var res int
	for _, c := range input {
		if c == '#' {
			res++
		}
	}
	return res
}

func (s *Solver) Part1(input string) any {
	sections := strings.Split(input, "\n\n")
	regions := sections[len(sections)-1]
	m := make(map[int]int)
	for i, section := range sections[:len(sections)-1] {
		m[i] = count(section)
	}
	var res int
	for _, region := range strings.Split(regions, "\n") {
		lst := strings.Split(region, ": ")
		first := lst[0]
		widelong := shared.Map(strconv.Atoi, strings.Split(first, "x"))
		area := widelong[0] * widelong[1]
		second := lst[1]
		total := 0
		for i, s := range strings.Split(second, " ") {
			n, _ := strconv.Atoi(s)
			total += n * m[i]
		}
		if total <= area {
			res++
		}
	}
	return res
}

func (s *Solver) Part2(input string) any {
	return 0
}

func main() {}
