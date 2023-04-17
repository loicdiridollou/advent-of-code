package main

import (
	"strings"
)

func Filter(vs []string, f func(string, int, string) bool, i int, element string) []string {
	vsm := make([]string, 0)
	for _, v := range vs {
		if f(v, i, element) {
			vsm = append(vsm, v)
		}
	}
	return vsm
}

func isElement(s string, i int, element string) bool {
	return string(s[i]) == element
}

func part2(input string) int {
	s := strings.Split(input, "\n")
	s = s[:len(s)-1]
	len_d := len(s[0])
	o2 := make([]string, len(s))
	co2 := make([]string, len(s))
	copy(o2, s)
	copy(co2, s)
	var o2_common string
	var co2_common string
	var o2_common_map map[string]int
	var co2_common_map map[string]int

	for i := 0; i < len_d; i++ {
		if len(co2) == 1 {
			continue
		}
		co2_common_map = Counter(Map(co2, extractElement, i))
		if co2_common_map["0"] <= co2_common_map["1"] {
			co2_common = "0"
		} else {
			co2_common = "1"
		}
		co2 = Filter(co2, isElement, i, co2_common)
	}
	for i := 0; i < len_d; i++ {
		if len(o2) == 1 {
			continue
		}
		o2_common_map = Counter(Map(o2, extractElement, i))
		if o2_common_map["0"] > o2_common_map["1"] {
			o2_common = "0"
		} else {
			o2_common = "1"
		}
		o2 = Filter(o2, isElement, i, o2_common)
	}

	return sliceToInt(strings.Split(o2[0], "")) * sliceToInt(strings.Split(co2[0], ""))
}
