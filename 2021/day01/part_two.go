package main

import (
	"strconv"
	"strings"
)

func strToInt(val string) int {
	a, _ := strconv.Atoi(val)
	return a
}

// part2 function
func part2(input string) int {
	s := strings.Split(input, "\n")
	res := 0
	sum := strToInt(s[0]) + strToInt(s[1]) + strToInt(s[2])
	var new_sum int

	for i := 3; i < len(s); i++ {
		new_sum = sum - strToInt(s[i-3]) + strToInt(s[i])
		if new_sum > sum {
			res++
		}
	}

	return res
}
