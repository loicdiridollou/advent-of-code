package main

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/mowshon/iterium"
)

func check_reflection_mul(p1 [2]int, p2 [2]int, val int, max_row int, max_col int) (Point, bool) {
	dr := p1[0] - p2[0]
	dc := p1[1] - p2[1]

	p3 := Point{p1[0] + val*dr, p1[1] + val*dc}

	if p3.r >= 0 && p3.r < max_row && p3.c >= 0 && p3.c < max_col {
		return p3, true
	}
	return p3, false
}

// part2 function
func part2(input string) int {
	lines := strings.Split(input, "\n")
	nodes := make(map[rune][][2]int)

	max_row := len(lines) - 1
	max_col := len(lines[0])

	for i, s := range lines {
		if s == "" {
			continue
		}
		for j, elem := range s {
			if unicode.IsLetter(elem) || unicode.IsDigit(elem) {
				_, exists := nodes[elem]
				if !exists {
					nodes[elem] = [][2]int{}
				}
				nodes[elem] = append(nodes[elem], [2]int{i, j})
			}
		}
	}

	reflections := make(map[string]bool)
	for _, node_list := range nodes {
		product, _ := iterium.Permutations(node_list, 2).Slice()
		for _, v1 := range product {
			idx := 0
			for {
				reflect, valid := check_reflection_mul(v1[0], v1[1], idx, max_row, max_col)
				if valid {
					reflections[strconv.Itoa(reflect.r)+"_"+strconv.Itoa(reflect.c)] = true
					idx += 1
				} else {
					break
				}
			}
		}
	}
	return len(reflections)
}
