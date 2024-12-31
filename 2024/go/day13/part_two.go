package main

import (
	"math"
	"strings"
)

// part2 function
func part2(input string) int {
	s := strings.Split(input, "\n\n")

	count := 0
	for _, block := range s {
		a, b := parse_block(block, 10000000000000).solve()
		if 0 < a && math.Floor(a) == a && 0 < b && math.Floor(b) == b {
			count += 3*int(a) + int(b)
		}
	}

	return count
}
