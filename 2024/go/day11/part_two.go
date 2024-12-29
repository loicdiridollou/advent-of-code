package main

import (
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func compute_pebbles_fast(num int, steps int) int {
	hashval := strconv.Itoa(num) + "_" + strconv.Itoa(steps)
	if val, ok := cache[hashval]; ok {
		return val
	}
	if steps == 0 {
		return 1
	}
	if num == 0 {
		res := compute_pebbles_fast(1, steps-1)
		hashval := strconv.Itoa(1) + "_" + strconv.Itoa(steps-1)
		cache[hashval] = res
		return res
	} else if str_num := strconv.Itoa(num); len(str_num)%2 == 0 {
		left, _ := strconv.Atoi(str_num[:len(str_num)/2])
		left_hashval := strconv.Itoa(left) + "_" + strconv.Itoa(steps-1)
		left_res := compute_pebbles_fast(left, steps-1)
		cache[left_hashval] = left_res

		right, _ := strconv.Atoi(str_num[len(str_num)/2:])
		right_hashval := strconv.Itoa(right) + "_" + strconv.Itoa(steps-1)
		right_res := compute_pebbles_fast(right, steps-1)
		cache[right_hashval] = right_res

		return left_res + right_res
	} else {
		res := compute_pebbles_fast(num*2024, steps-1)
		hashval := strconv.Itoa(num*2024) + "_" + strconv.Itoa(steps-1)
		cache[hashval] = res
		return res
	}
}

// part2 function
func part2(input string) int {
	s := strings.Split(input[:len(input)-1], " ")

	count := 0
	for _, pebble := range s {
		pebble_num, _ := strconv.Atoi(pebble)
		count += compute_pebbles_fast(pebble_num, 75)
	}

	return count
}
