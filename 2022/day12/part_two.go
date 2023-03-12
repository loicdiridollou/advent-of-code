package main

import "strings"

func part2(input string) int {
	s := strings.Split(input, "\n")
	res := make([][]string, 0)
	for _, u := range s[:len(s)-1] {
		spl := strings.Split(u, "")
		res = append(res, spl)
	}
	min_dist := len(res) * len(res[0])

	for i := range res {
		for j := range res[0] {
			if res[i][j] == "a" || res[i][j] == "S" {
				dist := BFS([2]int{i, j}, res)
				if dist != -1 && min_dist > dist {
					min_dist = dist
				}
			}
		}
	}
	return min_dist
}
