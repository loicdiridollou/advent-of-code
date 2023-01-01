package main

import (
	"os"
	"strings"
)

type Sensor struct {
	x int
	y int
	d int
}

func validPoint(x, y int, sensors []Sensor) bool {
	for _, sn := range sensors {
		dxy := Abs(x-sn.x) + Abs(y-sn.y)
		if dxy <= sn.d {
			return false
		}
	}
	return true
}

func part2() int {
	dat, _ := os.ReadFile("./day15-input")
	input := string(dat)[:len(string(dat))-1]
	s := strings.Split(input, "\n")

	var readings [][2][2]int
	for _, ln := range s {
		pts := parseLine(ln)
		readings = append(readings, [2][2]int{pts[0], pts[1]})
	}

	var sensors []Sensor
	for _, r := range readings {
		distance := Manhattan(r[0], r[1])
		sensors = append(sensors, Sensor{r[0][0], r[0][1], distance})
	}

	signs := [4][2]int{{-1, 1}, {-1, -1}, {1, -1}, {1, 1}}
	found := false

	for _, sn := range sensors {
		for dx := 0; dx < sn.d+2; dx++ {
			dy := (sn.d + 1) - dx
			for _, sign := range signs {
				x := sn.x + dx*sign[0]
				y := sn.y + dy*sign[1]
				if !(x >= 0 && x <= 4000000 && y >= 0 && y <= 4000000) {
					continue
				}
				if validPoint(x, y, sensors) && !found {
					found = true
					return x*4000000 + y
				}
			}
		}
	}
	return 0
}
