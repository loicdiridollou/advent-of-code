package main

import (
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func Map(vs []int, f func(int, int) int, args int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, args)
	}
	return vsm
}

func distance(v1, v2 int) int {
	if v1-v2 < 0 {
		return v2 - v1
	}
	return v1 - v2
}

func sum(input []int) int {
	res := 0
	for _, num := range input {
		res += num
	}
	return res
}

func prepareData(input string) []int {
	res := make([]int, 0)
	for _, num := range strings.Split(input[:len(input)-1], ",") {
		res = append(res, StrToInt(string(num)))
	}
	return res
}

func part1(input string) int {
	nums := prepareData(input)
	res := int(1e10)
	for i := 0; i < len(nums); i++ {
		if val := sum(Map(nums, distance, i)); val < res {
			res = val
		}
	}

	return res
}
