package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func processPair(input string) (int, int, int, int) {
	var pairs []string
	var p1, p2 []string
	pairs = strings.Split(input, ",")
	p1, p2 = strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
	return convInt(p1[0]), convInt(p1[1]), convInt(p2[0]), convInt(p2[1])
}

func checkContains(s1, e1, s2, e2 int) bool {
	if s1 <= s2 && e2 <= e1 {
		return true
	} else if s2 <= s1 && e1 <= e2 {
		return true
	} else {
		return false
	}
}

func checkOverlap(s1 int, e1 int, s2 int, e2 int) bool {
	if (s1 <= s2) && (s2 <= e1) {
		return true
	} else if (s2 <= s1) && (s1 <= e2) {
		return true
	} else {
		return false
	}
}

func part1() int {
	dat, _ := os.ReadFile("./day4-input")
	s := strings.Split(string(dat), "\n")
	var s1, e1, s2, e2 int
	var score int
	for i := 0; i < len(s)-1; i++ {
		s1, e1, s2, e2 = processPair(s[i])
		if checkContains(s1, e1, s2, e2) {
			score += 1
		}
	}
	return score
}

func part2() int {
	dat, _ := os.ReadFile("./day4-input")
	s := strings.Split(string(dat), "\n")
	var s1, e1, s2, e2 int
	var score int
	for i := 0; i < len(s)-1; i++ {
		s1, e1, s2, e2 = processPair(s[i])
		if checkOverlap(s1, e1, s2, e2) {
			score += 1
		}
	}
	return score
}

func main() {
	fmt.Println("Part 1 answer: ", part1())
	fmt.Println("Part 2 answer: ", part2())
	fmt.Println("DONE")
}
