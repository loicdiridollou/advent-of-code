package main

import (
	"fmt"
	"sort"
)

func combineMaps(map1 map[string]bool, map2 map[string]bool) map[string]bool {
	for k, v := range map2 {
		map1[k] = v
	}
	return map1
}

func isExplorable(grid [][]int, point Point, level int, visited map[string]bool) bool {
	_, pt_visited := visited[fmt.Sprint(point.row, "_", point.col)]
	in_grid := point.row >= 0 && point.row < len(grid) && point.col >= 0 && point.col < len(grid[0])
	if !in_grid {
		return false
	}
	is_flowing := grid[point.row][point.col] > level && grid[point.row][point.col] < 9
	return !pt_visited && in_grid && is_flowing
}

func findBasin(grid [][]int, point Point, visited map[string]bool) (int, map[string]bool) {
	// if we have already visited the point
	if visited[point.hash()] {
		return 0, visited
	}

	var num int
	var new_visited map[string]bool
	total_size := 1
	visited[point.hash()] = true
	level := grid[point.row][point.col]

	if isExplorable(grid, point.move("U"), level, visited) {
		num, new_visited = findBasin(grid, point.move("U"), visited)
		total_size += num
		visited = combineMaps(visited, new_visited)
	}

	if isExplorable(grid, point.move("D"), level, visited) {
		num, new_visited = findBasin(grid, point.move("D"), visited)
		total_size += num
		visited = combineMaps(visited, new_visited)
	}

	if isExplorable(grid, point.move("L"), level, visited) {
		num, new_visited = findBasin(grid, point.move("L"), visited)
		total_size += num
		visited = combineMaps(visited, new_visited)
	}

	if isExplorable(grid, point.move("R"), level, visited) {
		num, new_visited = findBasin(grid, point.move("R"), visited)
		total_size += num
		visited = combineMaps(visited, new_visited)
	}

	return total_size, visited
}

func part2(input string) int {
	grid := prepareData(input)
	basins := make([]int, 0)
	low_points := make([]Point, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if a, pt := isLowPoint(grid, i, j); a {
				low_points = append(low_points, pt)
			}
		}
	}

	for _, low_point := range low_points {
		size, _ := findBasin(grid, low_point, make(map[string]bool, 0))
		basins = append(basins, size)
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}
