package main

import "strings"

func part2(input string) int {
	s := strings.Split(input, "\n")
	var p1, p2, p3 string
	var score int
	for i := 0; i < (len(s)-1)/3; i++ {
		p1 = s[3*i]
		p2 = s[3*i+1]
		p3 = s[3*i+2]
		score += charToScore(findSameItemTriple(p1, p2, p3))
	}
	return score
}
