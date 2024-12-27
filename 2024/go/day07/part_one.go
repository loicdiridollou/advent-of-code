package main

import (
	"log"
	"strconv"
	"strings"
)

func search_value(target int, elements []int) bool {
	if len(elements) == 1 && elements[0] == target {
		return true
	} else if len(elements) == 1 && elements[0] != target {
		return false
	}

	return search_value(target, append([]int{elements[0] + elements[1]}, elements[2:]...)) ||
		search_value(target, append([]int{elements[0] * elements[1]}, elements[2:]...))
}

// part1 function
func part1(input string) int {
	s := strings.Split(input, "\n")

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		tmp := strings.Split(strings.ReplaceAll(s[i], ":", ""), " ")
		items := make([]int, 0, len(tmp))
		for _, raw := range tmp {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			items = append(items, v)
		}
		if search_value(items[0], items[1:]) {
			count += items[0]
		}
	}

	return count
}
