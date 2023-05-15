package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func BinToInt(input string) int {
	num, _ := strconv.ParseInt(input, 2, 64)
	return int(num)
}

func Map(vs string, f func(string) string) string {
	vsm := ""
	for _, v := range vs {
		vsm += f(string(v))
	}
	return vsm
}

func prepareData(input string) (string, map[string]bool) {
	val := strings.Split(input, "\n\n")
	points := make(map[string]bool, 0)

	for i, line := range strings.Split(val[1], "\n") {
		if len(line) == 0 {
			continue
		}
		for j, chr := range line {
			if string(chr) == "#" {
				points[fmt.Sprint(i, "_", j)] = true
			}
		}
	}
	return val[0], points
}

func flipOn(input string) string {
	if input == "#" {
		return "1"
	}
	return "0"
}

func getNeighbors(row, col int) [][2]int {
	neighbors := make([][2]int, 0)
	for drow := -1; drow < 2; drow++ {
		for dcol := -1; dcol < 2; dcol++ {
			neighbors = append(neighbors, [2]int{row + drow, col + dcol})
		}
	}
	return neighbors
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findDimensions(points map[string]bool) (int, int, int, int) {
	min_row, max_row, min_col, max_col := 10000, -10000, 10000, -10000

	for key := range points {
		pos := strings.Split(key, "_")
		row := StrToInt(pos[0])
		col := StrToInt(pos[1])
		min_row = MinInt(row, min_row)
		max_row = MaxInt(row, max_row)
		min_col = MinInt(col, min_col)
		max_col = MaxInt(col, max_col)
	}

	return min_row, max_row, min_col, max_col
}

func findNeighbors(row, col int, points map[string]bool) string {
	res := ""
	for _, point := range getNeighbors(row, col) {
		if _, ok := points[fmt.Sprint(point[0], "_", point[1])]; ok {
			res += "#"
		} else {
			res += "."
		}
	}

	return res
}

func plotImage(image map[string]bool) [][]string {
	res := make([][]string, 0)
	min_row, max_row, min_col, max_col := findDimensions(image)
	for row := min_row - 1; row <= max_row+1; row++ {
		tmp := make([]string, 0)
		for col := min_col - 1; col <= max_col+1; col++ {
			if _, ok := image[fmt.Sprint(row, "_", col)]; ok {
				tmp = append(tmp, "#")
			} else {
				tmp = append(tmp, ".")
			}
		}
		res = append(res, tmp)
	}

	return res
}

func padEdges(image map[string]bool, fill_value string) map[string]bool {
	if fill_value == "." {
		return image
	}

	// fill top row
	min_row, max_row, min_col, max_col := findDimensions(image)
	for i := min_col - 3; i <= max_col+3; i++ {
		image[fmt.Sprint(min_row-1, "_", i)] = true
		image[fmt.Sprint(min_row-2, "_", i)] = true
		image[fmt.Sprint(min_row-3, "_", i)] = true
		image[fmt.Sprint(max_row+1, "_", i)] = true
		image[fmt.Sprint(max_row+2, "_", i)] = true
		image[fmt.Sprint(max_row+3, "_", i)] = true
	}

	// fill first col
	for i := min_row - 3; i <= max_row+3; i++ {
		image[fmt.Sprint(i, "_", min_col-1)] = true
		image[fmt.Sprint(i, "_", min_col-2)] = true
		image[fmt.Sprint(i, "_", min_col-3)] = true
		image[fmt.Sprint(i, "_", max_row+1)] = true
		image[fmt.Sprint(i, "_", max_row+2)] = true
		image[fmt.Sprint(i, "_", max_row+3)] = true
	}

	return image
}

func enhanceImage(image map[string]bool, algo string, max_iter int) map[string]bool {
	fill_value := "."

	for iter := 0; iter < max_iter; iter++ {
		new_image := make(map[string]bool, 0)
		min_row, max_row, min_col, max_col := findDimensions(image)
		image = padEdges(image, fill_value)
		for row := min_row - 1; row <= max_row+1; row++ {
			for col := min_col - 1; col <= max_col+1; col++ {
				if string(algo[BinToInt(Map(findNeighbors(row, col, image), flipOn))]) == "#" {
					new_image[fmt.Sprint(row, "_", col)] = true
				}
			}
		}
		if fill_value == "." {
			fill_value = string(algo[0])
		} else {
			fill_value = string(algo[511])
		}
		image = new_image
	}
	return image
}

func part1(input string) int {
	algo, image := prepareData(input)

	return len(enhanceImage(image, algo, 2))
}
