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
	}
	return Point{}
}

func Map(vs []int, f func(int, int) int, args int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, args)
	}
	return vsm
}

func prepareData(input string) [][]int {
	res := make([][]int, 0)
	for _, num := range strings.Split(input, "\n") {
		if len(num) == 0 {
			continue
		}
		tmp_res := make([]int, 0)
		for _, i := range strings.Split(num, "") {
			tmp_res = append(tmp_res, StrToInt(string(i)))
		}
		res = append(res, tmp_res)
	}
	return res
}

func isLowPoint(grid [][]int, row_pos, col_pos int) (bool, Point) {
	points := make([]int, 0)

	if row_pos > 0 {
		points = append(points, grid[row_pos-1][col_pos])
	}
	if row_pos < len(grid)-1 {
		points = append(points, grid[row_pos+1][col_pos])
	}
	if col_pos > 0 {
		points = append(points, grid[row_pos][col_pos-1])
	}
	if col_pos < len(grid[0])-1 {
		points = append(points, grid[row_pos][col_pos+1])
	}

	for _, pt := range points {
		if pt <= grid[row_pos][col_pos] {
			return false, Point{}
		}
	}
	return true, Point{row_pos, col_pos}
}

func part1(input string) int {
	grid := prepareData(input)
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if a, _ := isLowPoint(grid, i, j); a {
				res += grid[i][j] + 1
			}
		}
	}

	return res
}
