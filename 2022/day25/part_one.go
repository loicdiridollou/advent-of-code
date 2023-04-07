package main

import (
	_ "embed"
	"strings"
)

func parseInput(input string) []string {
	res := make([]string, 0)
	for _, el := range strings.Split(input, "\n") {
		if len(el) == 0 {
			continue
		} else {
			res = append(res, string(el))
		}
	}
	return res
}

func MaxSize(lst ...string) int {
	max := len(lst[0])
	for _, el := range lst[1:] {
		if len(el) > max {
			max = len(el)
		}
	}
	return max
}

func convertChrToFloat(chr string) int {
	conv := map[string]int{"2": 2, "1": 1, "0": 0, "-": -1, "=": -2}
	return conv[chr]
}

func convertFloatToChr(num int) string {
	conv := map[int]string{2: "2", 1: "1", 0: "0", -1: "-", -2: "="}
	return conv[num]
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func part1(input string) string {
	numbers := parseInput(input)
	maxSize := MaxSize(numbers...)

	res := []string{}
	carry := 0

	for i := 0; i < maxSize; i++ {
		sum_col := carry
		for _, num := range numbers {
			if i < len(num) {
				sum_col += convertChrToFloat(string(num[len(num)-1-i]))
			}
		}
		carry = 0

		for sum_col > 2 {
			sum_col -= 5
			carry += 1
		}
		for sum_col < -2 {
			sum_col += 5
			carry -= 1
		}

		res = append(res, convertFloatToChr(sum_col))
	}

	return strings.Join(reverse(res), "")
}
