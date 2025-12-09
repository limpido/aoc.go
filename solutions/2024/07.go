package main

import (
	"strconv"
	"strings"

	"aoc.go/shared"
)

var Instance shared.Solver = &solution{}

type solution struct{}

func dfs(target, cur int, nums []int) bool {
	if len(nums) == 0 {
		return cur == target
	}
	return dfs(target, cur+nums[0], nums[1:]) || dfs(target, cur*nums[0], nums[1:])
}

func dfs2(target, cur int, nums []int) bool {
	if len(nums) == 0 {
		return cur == target
	}
	curStr := strconv.Itoa(cur)
	nStr := strconv.Itoa(nums[0])
	concatenated, _ := strconv.Atoi(curStr + nStr)
	return dfs2(target, cur+nums[0], nums[1:]) || dfs2(target, cur*nums[0], nums[1:]) || dfs2(target, concatenated, nums[1:])
}

func parseInput(input string) ([]int, [][]int) {
	lines := strings.Split(input, "\n")
	targets := make([]int, len(lines))
	arrs := make([][]int, len(lines))
	for i, line := range lines {
		l := strings.Split(line, ": ")
		targets[i], _ = strconv.Atoi(l[0])
		nums := strings.Split(l[1], " ")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			arrs[i] = append(arrs[i], n)
		}
	}
	return targets, arrs
}

func (s *solution) Part1(input string) any {
	var res int
	targets, arrs := parseInput(input)
	for i := range len(targets) {
		if dfs(targets[i], arrs[i][0], arrs[i][1:]) {
			res += targets[i]
		}
	}
	return res
}

func (s *solution) Part2(input string) any {
	var res int
	targets, arrs := parseInput(input)
	for i := range len(targets) {
		if dfs2(targets[i], arrs[i][0], arrs[i][1:]) {
			res += targets[i]
		}
	}
	return res
}

func main() {}
