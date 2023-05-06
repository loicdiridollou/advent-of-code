package main

import (
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func parseInput(lines []string) *Graph {
	g := &Graph{}
	g.s = len(lines)
	g.g = make([][]int, g.s)
	for i := 0; i < g.s; i++ {
		g.g[i] = make([]int, g.s)
	}

	for y, line := range lines {
		for x, cell := range line {
			g.g[y][x] = StrToInt(string(cell))
		}
	}

	return g
}

func part1(input string) int {
	lines := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	g := parseInput(lines)
	d, _ := g.dijkstra(0)
	return d[g.s*g.s]
}
