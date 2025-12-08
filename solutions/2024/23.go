package main

import (
	"fmt"
	"slices"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

type adjList map[string][]string

func parseInput(input string) adjList {
	lines := strings.Split(input, "\n")
	al := make(adjList)
	for _, line := range lines {
		tmp := strings.Split(line, "-")
		a, b := tmp[0], tmp[1]
		al[a] = append(al[a], b)
		al[b] = append(al[b], a)
	}
	return al
}

func (s *solution) Part1(input string) any {
	m := make(map[string][]string)
	al := parseInput(input)
	for n1 := range al {
		if n1[0] != 't' {
			continue
		}
		for _, n2 := range al[n1] {
			for _, n3 := range al[n2] {
				if slices.Contains(al[n1], n3) {
					lst := []string{n1, n2, n3}
					slices.Sort(lst)
					k := fmt.Sprintf("%s%s%s", lst[0], lst[1], lst[2])
					m[k] = lst
				}
			}
		}
	}
	return len(m)
}

func BronKerbosch(R, P, X []string, al adjList, maxClique *[]string) {
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(*maxClique) {
			*maxClique = slices.Clone(R)
		}
		return
	}

	pCopy := slices.Clone(P)
	for _, v := range pCopy {
		newR := append(R, v)
		neighbors := al[v]
		newP := shared.Intersect(P, neighbors)
		newX := shared.Intersect(X, neighbors)
		BronKerbosch(newR, newP, newX, al, maxClique)

		i := slices.Index(P, v)
		slices.Delete(P, i, i+1)
		X = append(X, v)
	}
}

func (s *solution) Part2(input string) any {
	var res string
	var maxClique []string
	al := parseInput(input)

	keys := make([]string, 0)
	for k := range al {
		keys = append(keys, k)
	}
	BronKerbosch([]string{}, keys, []string{}, al, &maxClique)
	slices.Sort(maxClique)
	for _, s := range maxClique {
		res += s + ","
	}
	return res
}

func main() {}
