package main

import (
	"regexp"
	"strconv"
)

type Area struct {
	x_min, x_max, y_min, y_max int
}

func (area *Area) contains(point Point) bool {
	if area.x_min <= point.x && point.x <= area.x_max && area.y_max >= point.y &&
		area.y_min <= point.y {
		return true
	}
	return false
}

func (area *Area) canReach(point Point) bool {
	if point.x <= area.x_max && point.y >= area.y_min {
		return true
	}
	return false
}

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

func prepareData(input string) Area {
	r, _ := regexp.Compile(`target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)`)
	matches := r.FindStringSubmatch(input)
	res := Map(matches[1:], StrToInt)
	return Area{res[0], res[1], res[2], res[3]}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Speed struct {
	x, y int
}

type Point struct {
	x, y int
}

func (point Point) update(speed Speed) Point {
	point.x += speed.x
	point.y += speed.y
	return point
}

func (speed Speed) update() Speed {
	speed.x = MaxInt(0, speed.x-1)
	speed.y--
	return speed
}

func part1(input string) int {
	area := prepareData(input)
	max_height := 0

	for i := 1; i < area.x_max; i++ {
		for j := area.y_min; j < 500; j++ {
			speed := Speed{i, j}
			point := Point{0, 0}
			max_y := area.y_min
			for area.canReach(point) {
				max_y = MaxInt(max_y, point.y)
				if area.contains(point) {
					max_height = MaxInt(max_height, max_y)
					break
				}
				point = point.update(speed)
				speed = speed.update()
			}
		}
	}

	return max_height
}
