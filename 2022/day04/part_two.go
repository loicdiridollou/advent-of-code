package main

import "strings"

func checkOverlap(s1 int, e1 int, s2 int, e2 int) bool {
	if (s1 <= s2) && (s2 <= e1) {
		return true
	} else if (s2 <= s1) && (s1 <= e2) {
		return true
	} else {
		return false
	}
}

func part2(input string) int {
	s := strings.Split(input, "\n")
	var s1, e1, s2, e2 int
	var score int
	for i := 0; i < len(s)-1; i++ {
		s1, e1, s2, e2 = processPair(s[i])
		if checkOverlap(s1, e1, s2, e2) {
			score += 1
		}
	}
	return score
}
