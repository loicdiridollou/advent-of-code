package main

import (
	"strings"
)

func (line *Line) isDiag() bool {
	p1 := line.p1
	p2 := line.p2
	if p1.x == p2.x {
		return false
	}
	num := abs(float64(p1.y-p2.y) / float64(p1.x-p2.x))
	return num == 1
}

func (line *Line) isOverlapDiag(line2 Line) []Point {
	dic_points := make(map[Point]bool, 0)
	res_points := make([]Point, 0)

	if (line.isAligned() && line2.isAligned()) || (line.isDiag() && line2.isAligned()) || (line.isAligned() && line2.isDiag()) || (line.isDiag() && line2.isDiag()) {
		for _, p := range line.enumerate() {
			dic_points[p] = true
		}
		for _, p := range line2.enumerate() {
			if dic_points[p] {
				res_points = append(res_points, p)
			} else {
				dic_points[p] = true
			}
		}
		return res_points
	} else {
		return []Point{}
	}
}

func part2(input string) int {
	s := strings.Split(input, "\n")

	lines := prepareData(s)
	dic_points := make(map[Point]bool, 0)
	res_points := make(map[Point]bool, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i].isAligned() || lines[i].isDiag() {
			for _, p := range lines[i].enumerate() {
				if dic_points[p] {
					res_points[p] = true
				} else {
					dic_points[p] = true
				}
			}
		}
	}
	return len(res_points)
}
