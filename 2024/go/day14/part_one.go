package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	r  int
	c  int
	dr int
	dc int
}

func euclideanRemainder(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

func (r Robot) move(action int, numr int, numc int) Robot {
	nr, nc := r.r+action*r.dr, r.c+action*r.dc
	nr, nc = euclideanRemainder(nr, numr), euclideanRemainder(nc, numc)
	return Robot{nr, nc, r.dr, r.dc}
}

func (r Robot) find_quadrant(numr int, numc int) int {
	if r.r < numr/2 && r.c < numc/2 {
		return 0
	} else if r.r < numr/2 && r.c >= numc/2+1 {
		return 1
	} else if r.r >= numr/2+1 && r.c >= numc/2+1 {
		return 2
	} else if r.r >= numr/2+1 && r.c < numc/2 {
		return 3
	}
	return -1
}

var re = regexp.MustCompile(`p=(.+),(.+) v=(.+),(.+)`)

func to_int(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func compute_score(quadrant_map map[int]int) int {
	count := 1
	for k, v := range quadrant_map {
		if k == -1 {
			continue
		}
		count *= v
	}
	return count
}

func parse_robot(line string) Robot {
	res := re.FindStringSubmatch(line)
	return Robot{to_int(res[2]), to_int(res[1]), to_int(res[4]), to_int(res[3])}
}

// part1 function
func part1(input string, numr int, numc int) int {
	s := strings.Split(input, "\n")

	quadrant_map := make(map[int]int)

	for _, line := range s {
		if line == "" {
			continue
		}

		quandrant := parse_robot(line).move(100, numr, numc).find_quadrant(numr, numc)
		if _, ok := quadrant_map[quandrant]; !ok {
			quadrant_map[quandrant] = 0
		}
		quadrant_map[quandrant]++
	}

	count := compute_score(quadrant_map)

	return count
}
