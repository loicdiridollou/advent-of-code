package main

import (
	_ "embed"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func parseLine(s string) [][2]int {
	r, _ := regexp.Compile("Sensor at x=(-*[0-9]+), y=(-*[0-9]+): closest beacon is at x=(-*[0-9]+), y=(-*[0-9]+)")
	matches := r.FindStringSubmatch(s)
	pts := make([][2]int, 2)
	pts[0][0], pts[0][1] = convInt(matches[1]), convInt(matches[2])
	pts[1][0], pts[1][1] = convInt(matches[3]), convInt(matches[4])
	return pts
}

type Interval struct {
	from, to int
}

type Intervals []Interval

func (is *Intervals) Add(from, to int) {
	// not the optimal implementation, but concise and fast enough for this case
	var l, r Intervals
	for _, interval := range *is {
		if interval.to < from {
			l = append(l, interval)
		} else if interval.from > to {
			r = append(r, interval)
		} else {
			if interval.from < from {
				from = interval.from
			}
			if interval.to > to {
				to = interval.to
			}
		}
	}
	*is = make(Intervals, len(l)+len(r)+1)
	copy(*is, l)
	(*is)[len(l)] = Interval{from, to}
	copy((*is)[len(l)+1:], r)
}

func part1() int {
	dat, _ := os.ReadFile("./day15-input")
	input := string(dat)[:len(string(dat))-1]
	row := 2000000
	s := strings.Split(input, "\n")

	var readings [][2][2]int
	for _, ln := range s {
		pts := parseLine(ln)
		readings = append(readings, [2][2]int{pts[0], pts[1]})
	}

	var intervals Intervals

	for _, r := range readings {
		distance := Manhattan(r[0], r[1])
		dx := distance - Abs(r[0][1]-row)
		if dx > 0 {
			from := r[0][0] - dx
			to := r[0][0] + dx
			if r[1][1] == row {
				switch r[1][0] {
				case from:
					from++
				case to:
					to--
				}
			}
			intervals.Add(from, to)
		}
	}

	count := 0
	for _, i := range intervals {
		count += i.to - i.from + 1
	}

	return count
}

func Manhattan(c1, c2 [2]int) int {
	return Abs(c1[1]-c2[1]) + Abs(c1[0]-c2[0])
}

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
