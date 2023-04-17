package main

import (
	"math"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func Map(vs []string, f func(string, int) string, args int) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, args)
	}
	return vsm
}

func extractElement(input string, i int) string {
	return string(input[i])
}

func Counter(s []string) map[string]int {
	dic := make(map[string]int)

	for _, el := range s {
		_, exists := dic[string(el)]
		if !exists {
			dic[string(el)] = 0
		}
		dic[string(el)] += 1
	}

	return dic
}

func mostCommon(dic map[string]int) string {
	var res string
	max := 0

	for key, el := range dic {
		if el >= max {
			max = el
			res = key
		}
	}

	return res
}

func leastCommon(dic map[string]int) string {
	var res string
	min := int(math.Inf(1))

	for key, el := range dic {
		if el <= min {
			min = el
			res = key
		}
	}

	return res
}

// MathPow calculates n to the mth power with the math.Pow() function
func MathPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}

func sliceToInt(s []string) int {
	res := 0
	max_pow := len(s) - 1

	for i := 0; i < len(s); i++ {
		if s[i] == "1" {
			res += MathPow(2, max_pow-i)
		}
	}
	return res
}

func part1(input string) int {
	s := strings.Split(input, "\n")
	s = s[:len(s)-1]
	// gamma_str := ""
	// epsilon_str := ""

	len_d := len(s[0])
	most_common := make([]string, 0)
	least_common := make([]string, 0)

	for i := 0; i < len_d; i++ {
		most_common = append(most_common, mostCommon(Counter(Map(s, extractElement, i))))
		least_common = append(least_common, leastCommon(Counter(Map(s, extractElement, i))))

	}

	return sliceToInt(most_common) * sliceToInt(least_common)
}
