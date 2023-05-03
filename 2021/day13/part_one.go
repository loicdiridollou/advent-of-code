package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

type Fold struct {
	axis string
	loc  int
}

func (point *Point) Hash() string {
	return fmt.Sprint(point.row, "_", point.col)
}

func (point *Point) fold(fld Fold) Point {
	if fld.axis == "y" {
		if point.row < fld.loc {
			return Point{point.row, point.col}
		} else {
			return Point{fld.loc - (point.row - fld.loc), point.col}
		}
	} else if fld.axis == "x" {
		if point.col < fld.loc {
			return Point{point.row, point.col}
		} else {
			return Point{point.row, fld.loc - (point.col - fld.loc)}
		}
	}
	return Point{}
}

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func prepareData(input string) ([]Point, []Fold) {
	split := strings.Split(input, "\n\n")
	points := make([]Point, 0)
	folds := make([]Fold, 0)

	for _, pt := range strings.Split(split[0], "\n") {
		if len(pt) == 0 {
			continue
		}
		coord := strings.Split(pt, ",")
		points = append(points, Point{StrToInt(coord[1]), StrToInt(coord[0])})
	}

	r, _ := regexp.Compile(`fold along (\w+)=(\d+)`)

	for _, fold := range strings.Split(split[1], "\n") {
		if len(fold) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(fold)
		folds = append(folds, Fold{matches[1], StrToInt(matches[2])})

	}
	return points, folds
}

func countUnique(points []Point) int {
	dic := make(map[string]bool, 0)
	res := 0

	for _, point := range points {
		if _, ok := dic[point.Hash()]; !ok {
			dic[point.Hash()] = true
			res++
		}
	}
	return res
}

func part1(input string) int {
	points, folds := prepareData(input)
	for _, fld := range folds[:1] {
		for i := 0; i < len(points); i++ {
			points[i] = points[i].fold(fld)
		}
	}

	return countUnique(points)
}
