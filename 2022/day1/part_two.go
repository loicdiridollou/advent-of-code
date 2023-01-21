package main

import (
	"os"
	"sort"
	"strings"
)

// part2 function
func part2() int {
	dat, _ := os.ReadFile("./day1-input")
	s := strings.Split(string(dat), "\n\n")
	var res []int

	for i := 0; i < len(s); i++ {
		res = append(res, addArray(strings.Split(s[i], "\n")))
	}
	sort.Ints(res)
	return addArrayInt(res[len(res)-3:])
}
