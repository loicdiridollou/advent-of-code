package main

import "fmt"

func plot(points []Point) [][]string {
	max_row, max_col := 0, 0

	for _, point := range points {
		if point.col > max_col {
			max_col = point.col
		}
		if point.row > max_row {
			max_row = point.row
		}
	}

	res := make([][]string, 0)
	for i := 0; i < max_row+1; i++ {
		tmp := make([]string, 0)
		for j := 0; j < max_col+1; j++ {
			tmp = append(tmp, ".")
		}
		res = append(res, tmp)
	}

	for _, point := range points {
		res[point.row][point.col] = "#"
	}

	return res
}

func part2(input string) int {
	points, folds := prepareData(input)
	for _, fld := range folds {
		for i := 0; i < len(points); i++ {
			points[i] = points[i].fold(fld)
		}
	}

	fmt.Println(plot(points))

	return countUnique(points)
}
