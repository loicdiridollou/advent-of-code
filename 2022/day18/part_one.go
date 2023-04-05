package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

type Point struct {
	x, y, z int
}

func parseLine(input string) []Point {
	s := strings.Split(string(input), "\n")
	var list []Point

	for _, el := range s {
		if len(el) == 0 {
			continue
		}
		coord := strings.Split(el, ",")
		list = append(list, Point{convInt(coord[0]), convInt(coord[1]), convInt(coord[2])})
	}
	return list
}

func hashVal(x, y, z int) string {
	return fmt.Sprint(x) + "_" + fmt.Sprint(y) + "_" + fmt.Sprint(z)
}

func DFS(point Point, points []Point, points_map map[string]Point, visited map[string]bool) (int, map[string]bool) {
	faces := 0
	x, y, z := point.x, point.y, point.z
	act := hashVal(x, y, z)
	up := hashVal(x, y, z+1)
	dn := hashVal(x, y, z-1)
	lf := hashVal(x-1, y, z)
	rt := hashVal(x+1, y, z)
	ft := hashVal(x, y+1, z)
	re := hashVal(x, y-1, z)

	visited[act] = true
	var ok, vst bool
	var fcs int

	if _, ok = points_map[up]; ok {
		if _, vst = visited[up]; !vst {
			fcs, visited = DFS(Point{x, y, z + 1}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}
	if _, ok = points_map[dn]; ok {
		if _, vst = visited[dn]; !vst {
			fcs, visited = DFS(Point{x, y, z - 1}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}
	if _, ok = points_map[lf]; ok {
		if _, vst = visited[lf]; !vst {
			fcs, visited = DFS(Point{x - 1, y, z}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}
	if _, ok = points_map[rt]; ok {
		if _, vst = visited[rt]; !vst {
			fcs, visited = DFS(Point{x + 1, y, z}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}
	if _, ok = points_map[ft]; ok {
		if _, vst = visited[ft]; !vst {
			fcs, visited = DFS(Point{x, y + 1, z}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}
	if _, ok = points_map[re]; ok {
		if _, vst = visited[re]; !vst {
			fcs, visited = DFS(Point{x, y - 1, z}, points, points_map, visited)
			faces += fcs
		}
	} else {
		faces++
	}

	return faces, visited
}

func part1(input string) int {
	points := parseLine(input)

	points_map := make(map[string]Point, 0)
	for _, point := range points {
		points_map[hashVal(point.x, point.y, point.z)] = point
	}
	visited := make(map[string]bool, 0)

	faces := 0
	var fcs int
	for _, point := range points {
		if _, vst := visited[hashVal(point.x, point.y, point.z)]; !vst {
			fcs, visited = DFS(point, points, points_map, visited)
			faces += fcs
		}
	}

	return faces
}
