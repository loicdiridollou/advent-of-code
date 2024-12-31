package main

import (
	"strconv"
	"strings"
)

func hash_pos(r int, c int) string {
	return strconv.Itoa(r) + "_" + strconv.Itoa(c)
}

func dfs(
	grid [][]string,
	r int,
	c int,
	visited map[string]bool,
) (int, int) {
	area := 1
	perimeter := 0
	visited[hash_pos(r, c)] = true
	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nr := r + dir[0]
		nc := c + dir[1]
		_, ok := visited[hash_pos(nr, nc)]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) &&
			grid[nr][nc] == grid[r][c] && !ok {
			new_area, new_perimeter := dfs(grid, nr, nc, visited)
			area += new_area
			perimeter += new_perimeter
		} else if !ok || (ok && grid[nr][nc] != grid[r][c]) {
			perimeter += 1
		}
	}
	return area, perimeter
}

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n")

	grid := [][]string{}

	for r, row := range s {
		if row == "" {
			continue
		}
		grid = append(grid, []string{})
		for _, col := range row {
			grid[r] = append(grid[r], string(col))
		}
	}

	visited := make(map[string]bool)
	count := 0
	for r, row := range s {
		if row == "" {
			continue
		}
		for c := range row {
			if val, ok := visited[hash_pos(r, c)]; ok && val {
				continue
			}
			area, perimeter := dfs(grid, r, c, visited)
			count += area * perimeter
		}
	}

	return count
}
