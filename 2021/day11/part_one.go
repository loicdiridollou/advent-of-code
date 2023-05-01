package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

type Point struct {
	row, col int
}

func (point *Point) hash() string {
	return fmt.Sprint(point.row, "_", point.col)
}

func (point *Point) move(direction string) Point {
	switch direction {
	case "U":
		return Point{point.row - 1, point.col}
	case "D":
		return Point{point.row + 1, point.col}
	case "L":
		return Point{point.row, point.col - 1}
	case "R":
		return Point{point.row, point.col + 1}
	case "UL":
		return Point{point.row - 1, point.col - 1}
	case "UR":
		return Point{point.row - 1, point.col + 1}
	case "DL":
		return Point{point.row + 1, point.col - 1}
	case "DR":
		return Point{point.row + 1, point.col + 1}
	}
	return Point{}
}

func (point *Point) isValid(row_dim, col_dim int) bool {
	if point.row < 0 || point.row >= row_dim {
		return false
	}
	if point.col < 0 || point.col >= col_dim {
		return false
	}
	return true
}

func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func prepareData(input string) [][]int {
	res := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		tmp_res := make([]int, 0)
		for _, chr := range line {
			tmp_res = append(tmp_res, StrToInt(string(chr)))
		}
		res = append(res, tmp_res)
	}
	return res
}

func addOneInt(val int) int {
	return val + 1
}

func addOne(grid [][]int) [][]int {
	for i := 0; i < len(grid); i++ {
		grid[i] = Map(grid[i], addOneInt)
	}
	return grid
}

func getNeighbors(grid [][]int, point Point) []Point {
	list_points := make([]Point, 0)

	for _, dir := range []string{"UL", "U", "UR", "L", "R", "DL", "D", "DR"} {
		if new_point := point.move(dir); new_point.isValid(len(grid), len(grid[0])) {
			list_points = append(list_points, new_point)
		}
	}
	return list_points
}

func increaseEnergy(grid [][]int, list_points []Point) [][]int {
	for _, point := range list_points {
		grid[point.row][point.col]++
	}
	return grid
}

func flashOctopus(grid [][]int) ([][]int, int) {
	num_flash := 0
	flashed_points := make(map[string]bool, 0)
	num_changes := 1

	for num_changes > 0 {
		num_changes = 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				_, flashed := flashed_points[fmt.Sprint(i, "_", j)]
				if grid[i][j] > 9 && !flashed {
					neighbors := getNeighbors(grid, Point{i, j})
					grid = increaseEnergy(grid, neighbors)
					flashed_points[fmt.Sprint(i, "_", j)] = true
					num_changes++
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 9 {
				grid[i][j] = 0
				num_flash++
			}
		}
	}
	return grid, num_flash
}

func part1(input string) int {
	grid := prepareData(input)
	num_rounds := 100
	total_flashes := 0
	var num_flashes int

	for i := 0; i < num_rounds; i++ {
		grid = addOne(grid)
		grid, num_flashes = flashOctopus(grid)
		total_flashes += num_flashes
	}

	return total_flashes
}
