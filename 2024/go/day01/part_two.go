package main

import (
	"strconv"
	"strings"
)

func counter(words []int) map[int]int {
	counter := make(map[int]int)

	for _, word := range words {
		counter[word]++
	}

	return counter
}

// part2 function
func part2(input string) int {
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

	counter_map := counter(v2)

	diff := 0
	for _, val := range v1 {
		diff += val * counter_map[val]
	}
	return diff
}
