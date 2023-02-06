package main

import (
	"sort"
	"strings"
)

// part2 function
func part2(input string) int {
	s := strings.Split(input, "\n\n")
	var res []int

	for i := 0; i < len(s); i++ {
		res = append(res, addArray(strings.Split(s[i], "\n")))
	}
	sort.Ints(res)
	return addArrayInt(res[len(res)-3:])
}
