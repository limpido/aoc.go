package main

import (
	"sort"
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func parseInput(input string) []shared.Point3D {
	lines := strings.Split(input, "\n")
	arr := make([]shared.Point3D, 0)
	for _, l := range lines {
		m := strings.Split(l, ",")
		n := shared.Map(strconv.Atoi, m)
		p := shared.Point3D{n[0], n[1], n[2]}
		arr = append(arr, p)
	}
	return arr
}

type Edge struct {
	dist, p1, p2 int
}

func (s *solution) Part1(input string) any {
	var res int
	points := parseInput(input)

	edges := make([]Edge, 0)
	for i := range points {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			d := shared.EuclideanDistance3D(p1, p2)
			e := Edge{d, i, j}
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	uf := shared.NewUnionFind(len(points))
	for i := range points {
		uf.Add(i)
	}

	// make the n shortest connections
	n := 1000
	for _ = range n {
		e := edges[0]
		edges = edges[1:]
		uf.Union(e.p1, e.p2)
	}

	sort.Slice(uf.Size, func(i, j int) bool {
		return uf.Size[i] > uf.Size[j]
	})
	res = uf.Size[0] * uf.Size[1] * uf.Size[2]
	return res
}

func (s *solution) Part2(input string) any {
	var res int
	points := parseInput(input)

	edges := make([]Edge, 0)
	for i := range points {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			d := shared.EuclideanDistance3D(p1, p2)
			e := Edge{d, i, j}
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	uf := shared.NewUnionFind(len(points))
	for i := range points {
		uf.Add(i)
	}

	// keep connecting until all in the same circuit
	for {
		e := edges[0]
		edges = edges[1:]
		uf.Union(e.p1, e.p2)
		if uf.Size[uf.Find(e.p1)] == len(points) {
			res = points[e.p1].X * points[e.p2].X
			break
		}
	}
	return res
}

func main() {}
