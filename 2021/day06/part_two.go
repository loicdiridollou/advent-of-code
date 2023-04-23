package main

import (
	"strings"
)

func sum(input []int) int {
	res := 0
	for _, num := range input {
		res += num
	}
	return res
}

func part2(input string) int {
	s := strings.Split(input[:len(input)-1], ",")
	vs := Map(s, StrToInt)
	nums := make([]int, 9)
	for i := 0; i < len(s); i++ {
		nums[vs[i]]++
	}
	new_fish := 0
	for j := 0; j < 256; j++ {
		for i := 0; i < 9; i++ {
			if i == 0 {
				new_fish = nums[0]
			} else {
				nums[i-1] = nums[i]
			}
		}
		nums[8] = new_fish
		nums[6] += new_fish
	}

	return sum(nums)
}
