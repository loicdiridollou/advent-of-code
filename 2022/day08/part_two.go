package main

import "strings"

func part2(input string) int {
	s := strings.Split(input, "\n")
	tbl := formatInput(s)
	totVisible := make([]int, 0)

	for i := 1; i < len(tbl)-1; i++ {
		for j := 1; j < len(tbl[0])-1; j++ {
			totVisible = append(totVisible, computeScore(i, j, tbl))
		}
	}
	return MaxArray(totVisible)
}
