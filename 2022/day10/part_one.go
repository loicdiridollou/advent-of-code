package main

import (
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func part1(input string) int {
	s := strings.Split(input, "\n")
	strength := []int{1}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == "noop" {
			strength = append(strength, strength[len(strength)-1])
		} else {
			toAdd := convInt(strings.Split(s[i], " ")[1])
			strength = append(strength, strength[len(strength)-1])
			strength = append(strength, strength[len(strength)-1]+toAdd)
		}
	}
	return 20*strength[19] + 60*strength[59] + 100*strength[99] + 140*strength[139] +
		180*strength[179] + 220*strength[219]
}
