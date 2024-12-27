package main

import (
	"log"
	"strconv"
	"strings"
)

func search_value_new(target int, elements []int) bool {
	if len(elements) == 1 && elements[0] == target {
		return true
	} else if len(elements) == 1 && elements[0] != target {
		return false
	}
	new_value0 := strconv.Itoa(elements[0])
	new_value1 := strconv.Itoa(elements[1])
	new_val, _ := strconv.Atoi(new_value0 + new_value1)

	return search_value_new(target, append([]int{elements[0] + elements[1]}, elements[2:]...)) ||
		search_value_new(target, append([]int{elements[0] * elements[1]}, elements[2:]...)) ||
		search_value_new(target, append([]int{new_val}, elements[2:]...))
}

// part2 function
func part2(input string) int {
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
		if search_value_new(items[0], items[1:]) {
			count += items[0]
		}
	}

	return count
}
