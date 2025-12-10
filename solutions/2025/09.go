package main

import (
	"aoc.go/shared"
	"fmt"
	"strconv"
	"strings"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseInput(input string) [][]int {
	tiles := [][]int{}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		n := shared.Map(strconv.Atoi, strings.Split(l, ","))
		x, y := n[0], n[1]
		tiles = append(tiles, []int{x, y})
	}
	return tiles
}

func (s *Solver) Part1(input string) any {
	var res int
	tiles := parseInput(input)
	for i, tile1 := range tiles {
		x1, y1 := tile1[0], tile1[1]
		for j := i + 1; j < len(tiles); j++ {
			tile2 := tiles[j]
			x2, y2 := tile2[0], tile2[1]
			area := (shared.Abs(x1-x2) + 1) * (shared.Abs(y1-y2) + 1)
			res = max(res, area)
		}
	}
	return res
}

func (s *Solver) Part2(input string) any {
	return
}

func main() {}
