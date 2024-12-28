package main

import (
	"strconv"
	"strings"
)

// part2 function
func part2(input string) int {
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
				count += len(search_trail(grid, r, c))
			}
		}
	}

	return count
}
