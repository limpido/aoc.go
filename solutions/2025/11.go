package main

import (
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &Solver{}

type Solver struct{}

func parseInput(input string) map[string][]string {
	lines := strings.Split(input, "\n")
	m := map[string][]string{}
	for _, line := range lines {
		kv := strings.Split(line, ": ")
		k := kv[0]
		m[k] = strings.Split(kv[1], " ")
	}
	return m
}

func pathCount(inp string, cache map[string]int, adj map[string][]string) int {
	if inp == "out" {
		return 1
	}
	if v, ok := cache[inp]; ok {
		return v
	}
	var res int
	for _, nb := range adj[inp] {
		res += pathCount(nb, cache, adj)
	}
	cache[inp] = res
	return res
}

func (s *Solver) Part1(input string) any {
	adj := parseInput(input)
	cache := map[string]int{}
	return pathCount("you", cache, adj)
}

const (
	SEEN_FFT = 1
	SEEN_DAC = 2
)

type Key struct {
	input string
	mask  int
}

func pathCount2(inp string, mask int, cache map[Key]int, adj map[string][]string) int {
	if inp == "out" {
		if mask == (SEEN_FFT | SEEN_DAC) {
			return 1
		}
		return 0
	}
	key := Key{inp, mask}
	if v, ok := cache[key]; ok {
		return v
	}
	var res int
	for _, nb := range adj[inp] {
		newMask := mask
		if nb == "fft" {
			newMask |= SEEN_FFT
		}
		if nb == "dac" {
			newMask |= SEEN_DAC
		}
		res += pathCount2(nb, newMask, cache, adj)
	}
	cache[key] = res
	return res
}

func (s *Solver) Part2(input string) any {
	adj := parseInput(input)
	cache := make(map[Key]int)
	return pathCount2("svr", 0, cache, adj)
}

func main() {}
