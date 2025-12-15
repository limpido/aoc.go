package main

import (
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseInput(input string) []string {
	return strings.Split(input, " ")
}

func parseInt(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		return s[i:]
	}
	return "0"
}

func (s *Solver) Part1(input string) any {
	stones := parseInput(input)
	cur := stones
	for i := 0; i < 25; i++ {
		var next []string
		for _, st := range cur {
			if st == "0" {
				next = append(next, "1")
				continue
			}
			if len(st)%2 == 0 {
				size := len(st)
				next = append(next, st[:size/2])
				next = append(next, parseInt(st[size/2:]))
				continue
			}
			n, _ := strconv.Atoi(st)
			s := strconv.Itoa(n * 2024)
			next = append(next, s)
		}
		cur = next

	}
	return len(cur)
}

type MemoKey struct {
	stone  string
	blinks int
}

var memo = make(map[MemoKey]int)

func count(stone string, blinks int) int {
	if blinks == 0 {
		return 1
	}
	k := MemoKey{stone, blinks}
	if res, ok := memo[k]; ok {
		return res
	}
	var res int
	if stone == "0" {
		res = count("1", blinks-1)
	} else if len(stone)%2 == 0 {
		size := len(stone)
		left := stone[:size/2]
		right := parseInt(stone[size/2:])
		res = count(left, blinks-1) + count(right, blinks-1)
	} else {
		n, _ := strconv.Atoi(stone)
		res = count(strconv.Itoa(n*2024), blinks-1)
	}
	memo[k] = res
	return res
}

func (s *Solver) Part2(input string) any {
	stones := parseInput(input)
	res := 0
	k := MemoKey{"0", 1}
	memo[k] = 1
	for _, st := range stones {
		res += count(st, 75)
	}
	return res
}

func main() {}
