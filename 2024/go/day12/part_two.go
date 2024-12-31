package main

import (
	"strings"
)

func dfs_with_corners(
	grid [][]string,
	r int,
	c int,
	visited map[string]bool,
) (int, int) {
	area := 1
	corners := 0
	visited[hash_pos(r, c)] = true

	directions := [][2][2]int{
		{{0, 1}, {1, 0}},
		{{1, 0}, {0, -1}},
		{{0, -1}, {-1, 0}},
		{{-1, 0}, {0, 1}},
	}

	for _, pair := range directions {
		r1, c1, r2, c2 := pair[0][0], pair[0][1], pair[1][0], pair[1][1]
		nr, nc := r+r1, c+c1
		test_a := (0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] == grid[r][c])

		nr, nc = r+r2, c+c2
		test_b := (0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] == grid[r][c])

		nr, nc = r+r1+r2, c+c1+c2
		test_c := (0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) && grid[nr][nc] != grid[r][c])

		if test_a && test_b && test_c {
			corners += 1
		} else if !test_a && !test_b {
			corners += 1
		}
	}

	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nr := r + dir[0]
		nc := c + dir[1]
		_, ok := visited[hash_pos(nr, nc)]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) &&
			grid[nr][nc] == grid[r][c] && !ok {
			new_area, new_corners := dfs_with_corners(grid, nr, nc, visited)
			area += new_area
			corners += new_corners
		}
	}
	return area, corners
}

// part2 function
func part2(input string) int {
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
			area, corners := dfs_with_corners(grid, r, c, visited)
			count += area * corners
		}
	}

	return count
}
