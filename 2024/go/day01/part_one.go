package main

import (
	"sort"
	"strconv"
	"strings"
)

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n")
	var v1 []int
	var v2 []int

	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		spt := delete_empty(strings.Split(s[i], " "))

		a1, _ := strconv.Atoi(spt[0])
		a2, _ := strconv.Atoi(spt[1])
		v1 = append(v1, a1)
		v2 = append(v2, a2)
	}
	sort.Ints(v1)
	sort.Ints(v2)

	diff := 0
	for i := range v1 {
		diff += abs(v1[i] - v2[i])
	}

	return diff
}
