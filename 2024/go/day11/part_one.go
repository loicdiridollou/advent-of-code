package main

import (
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

func compute_pebbles(num int, steps int) int {
	if steps == 0 {
		return 1
	}
	if num == 0 {
		return compute_pebbles(1, steps-1)
	} else if str_num := strconv.Itoa(num); len(str_num)%2 == 0 {
		left, _ := strconv.Atoi(str_num[:len(str_num)/2])
		right, _ := strconv.Atoi(str_num[len(str_num)/2:])
		return compute_pebbles(left, steps-1) + compute_pebbles(right, steps-1)
	} else {
		return compute_pebbles(num*2024, steps-1)
	}
}

// part1 function
func part1(input string) int {
	s := strings.Split(input[:len(input)-1], " ")

	count := 0
	for _, pebble := range s {
		pebble_num, _ := strconv.Atoi(pebble)
		count += compute_pebbles(pebble_num, 25)
	}

	return count
}
