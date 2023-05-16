package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

type Operation struct {
	type_op                                  string
	min_x, max_x, min_y, max_y, min_z, max_z int
}

func prepareData(input string) []Operation {
	res := make([]Operation, 0)
	r, _ := regexp.Compile(`(\w+) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		res = append(res,
			Operation{
				matches[1],
				StrToInt(matches[2]),
				StrToInt(matches[3]),
				StrToInt(matches[4]),
				StrToInt(matches[5]),
				StrToInt(matches[6]),
				StrToInt(matches[7]),
			},
		)
	}
	return res
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1(input string) int {
	operations := prepareData(input)
	dic_cubes := make(map[string]bool, 0)

	for _, operation := range operations {
		for i := MaxInt(-50, operation.min_x); i <= MinInt(50, operation.max_x); i++ {
			for j := MaxInt(-50, operation.min_y); j <= MinInt(50, operation.max_y); j++ {
				for k := MaxInt(-50, operation.min_z); k <= MinInt(50, operation.max_z); k++ {
					if operation.type_op == "on" {
						dic_cubes[fmt.Sprint(i, "_", j, "_", k)] = true
					} else {
						delete(dic_cubes, fmt.Sprint(i, "_", j, "_", k))
					}
				}
			}
		}
	}
	return len(dic_cubes)
}
