package main

import (
	"strconv"
	"strings"
)

func search_trail(grid [][]int, r int, c int) []string {
	if grid[r][c] == 9 {
		return []string{strconv.Itoa(r) + "_" + strconv.Itoa(c)}
	}
	res := []string{}

	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nr := r + dir[0]
		nc := c + dir[1]
		if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) &&
			grid[nr][nc] == grid[r][c]+1 {
			res = append(res, search_trail(grid, nr, nc)...)
		}
	}

	return res
}

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n")
	grid := [][]int{}

	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		tmp := []int{}
		for _, el := range s[i] {
			new_val, _ := strconv.Atoi(string(el))
			tmp = append(tmp, new_val)
		}
		grid = append(grid, tmp)
	}

	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				hashMap := make(map[string]int)

				for i, val := range search_trail(grid, r, c) {
					hashMap[val] = i
				}
				count += len(hashMap)
			}
		}
	}

	return count
}
