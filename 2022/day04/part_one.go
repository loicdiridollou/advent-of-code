package main

import (
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func processPair(input string) (int, int, int, int) {
	var pairs []string
	var p1, p2 []string
	pairs = strings.Split(input, ",")
	p1, p2 = strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
	return convInt(p1[0]), convInt(p1[1]), convInt(p2[0]), convInt(p2[1])
}

func checkContains(s1, e1, s2, e2 int) bool {
	if s1 <= s2 && e2 <= e1 {
		return true
	} else if s2 <= s1 && e1 <= e2 {
		return true
	} else {
		return false
	}
}

func part1(input string) int {
	s := strings.Split(input, "\n")
	var s1, e1, s2, e2 int
	var score int
	for i := 0; i < len(s)-1; i++ {
		s1, e1, s2, e2 = processPair(s[i])
		if checkContains(s1, e1, s2, e2) {
			score += 1
		}
	}
	return score
}
