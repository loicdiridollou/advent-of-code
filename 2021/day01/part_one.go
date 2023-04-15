package main

import (
	"strconv"
	"strings"
)

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n")
	res := 0

	for i := 0; i < len(s)-1; i++ {
		a, _ := strconv.Atoi(s[i])
		b, _ := strconv.Atoi(s[i+1])
		if a < b {
			res++
		}
	}

	return res
}
