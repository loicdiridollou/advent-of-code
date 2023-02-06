package main

import (
	"strconv"
	"strings"
)

func addArray(numbs []string) int {
	result := 0
	for _, numb := range numbs {
		num, _ := strconv.Atoi(numb)
		result += num
	}
	return result
}

func addArrayInt(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}

func maxArray(numbs []int) int {
	for j := 1; j < len(numbs); j++ {
		if numbs[j] > numbs[0] {
			numbs[0] = numbs[j]
		}
	}
	return numbs[0]
}

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n\n")
	var res []int

	for i := 0; i < len(s); i++ {
		res = append(res, addArray(strings.Split(s[i], "\n")))
	}
	return maxArray(res)
}
