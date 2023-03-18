package main

import "strings"

func part2(input string) int {
	s := strings.Split(input[:len(input)-1], "\n")
	vertices := make([][][2]int, 0)
	m1, m2, h := 0, 0, 0
	var tmp [][2]int
	for _, st := range s {
		tmp = make([][2]int, 0)
		for _, pts := range strings.Split(st, " -> ") {
			coord := strings.Split(pts, ",")
			p1, p2 := convInt(string(coord[0])), convInt(string(coord[1]))
			if p1 > m2 {
				m2 = p1
			}
			if p2 > h {
				h = p2
			}
			tmp = append(tmp, [2]int{p1, p2})
		}
		vertices = append(vertices, tmp)
	}

	m2 += 500
	plan := make([][]string, h+3)
	for i := 0; i < len(plan); i++ {
		plan[i] = make([]string, m2-m1+1)
		for j := 0; j < len(plan[0]); j++ {
			plan[i][j] = "."
		}
	}

	for _, v := range vertices {
		for i := 0; i < len(v)-1; i++ {
			p1, p2 := v[i], v[i+1]
			if p1[0] == p2[0] {
				if p1[1] > p2[1] {
					p1, p2 = p2, p1
				}
				for k := p1[1]; k <= p2[1]; k++ {
					plan[k][p1[0]-m1] = "#"
				}
			} else {
				if p1[0] > p2[0] {
					p1, p2 = p2, p1
				}
				for l := p1[0] - m1; l <= p2[0]-m1; l++ {
					plan[p1[1]][l] = "#"
				}
			}
		}
	}
	for i := 0; i < len(plan[0]); i++ {
		plan[len(plan)-1][i] = "#"
	}

	out := false
	var unit int
	for !out {
		ctn := true
		pos := [2]int{0, 500 - m1}
		for ctn && !out {
			if pos[0]+1 == len(plan) {
				ctn = false
				out = true
			} else if plan[pos[0]+1][pos[1]] == "." {
				pos[0]++
			} else if pos[1]-1 >= 0 && plan[pos[0]+1][pos[1]-1] == "." {
				pos[0]++
				pos[1]--
			} else if pos[1]+1 < len(plan[0]) && plan[pos[0]+1][pos[1]+1] == "." {
				pos[0]++
				pos[1]++
			} else if 0 < pos[0] && pos[0] <= len(plan) && 0 < pos[1] && pos[1] < len(plan[0]) {
				plan[pos[0]][pos[1]] = "o"
				ctn = false
			} else {
				ctn = false
				out = true
			}
		}
		unit++
	}
	return unit
}
