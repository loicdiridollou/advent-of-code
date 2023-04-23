package main

import (
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func Map(vs []string, f func(string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

type Line struct {
	p1, p2 Point
}

type Point struct {
	x, y int
}

func (line *Line) isAligned() bool {
	p1 := line.p1
	p2 := line.p2
	return p1.x == p2.x || p1.y == p2.y
}

func abs(x float64) int {
	if x < 0 {
		return int(-x)
	} else {
		return int(x)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (line *Line) enumerate() []Point {
	res := make([]Point, 0)
	if line.p1.x == line.p2.x {
		for i := min(line.p1.y, line.p2.y); i <= max(line.p1.y, line.p2.y); i++ {
			res = append(res, Point{line.p1.x, i})
		}
	} else if line.p1.y == line.p2.y {
		for i := min(line.p1.x, line.p2.x); i <= max(line.p1.x, line.p2.x); i++ {
			res = append(res, Point{i, line.p1.y})
		}
	} else if line.isDiag() {
		var y_sgn, x_sgn int
		if line.p1.x > line.p2.x {
			x_sgn = -1
		} else {
			x_sgn = 1
		}
		if line.p1.y > line.p2.y {
			y_sgn = -1
		} else {
			y_sgn = 1
		}
		for i := 0; i <= max(line.p1.x-line.p2.x, line.p2.x-line.p1.x); i++ {
			res = append(res, Point{line.p1.x + i*x_sgn, line.p1.y + i*y_sgn})
		}
	}
	return res
}

func (line *Line) isOverlap(line2 Line) []Point {
	dic_points := make(map[Point]bool, 0)
	res_points := make([]Point, 0)

	if line.isAligned() && line2.isAligned() {
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

func prepareData(input []string) []Line {
	r, _ := regexp.Compile(`(\d+),(\d+) -> (\d+),(\d+)`)
	lines := make([]Line, 0)
	for _, grp := range input {
		for _, line := range strings.Split(grp, "\n") {
			if len(line) == 0 {
				continue
			}
			matches := r.FindStringSubmatch(line)
			points := Map(matches[1:], StrToInt)
			line := Line{Point{points[0], points[1]}, Point{points[2], points[3]}}
			lines = append(lines, line)
		}
	}
	return lines
}

func part1(input string) int {
	s := strings.Split(input, "\n")

	lines := prepareData(s)
	dic := make(map[Point]bool, 0)
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			for _, p := range lines[i].isOverlap(lines[j]) {
				dic[p] = true
			}
		}
	}
	return len(dic)
}
