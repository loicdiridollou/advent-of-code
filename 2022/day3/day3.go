package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func findSameItem(p1 string, p2 string) string {
	set := make(map[string]bool)
	for i := 0; i < len(p1); i++ {
		set[string(p1[i])] = true
	}
	var res string
	for i := 0; i < len(p2); i++ {
		if set[string(p2[i])] {
			res = string(p2[i])
		}
	}
	return res
}

func findSameItemTriple(p1 string, p2 string, p3 string) string {
	set1 := make(map[string]bool)
	set := make(map[string]bool)
	for i := 0; i < len(p1); i++ {
		set1[string(p1[i])] = true
	}
	for i := 0; i < len(p2); i++ {
		if set1[string(p2[i])] {
			set[string(p2[i])] = true
		}
	}
	var res string
	for i := 0; i < len(p3); i++ {
		if set[string(p3[i])] {
			res = string(p3[i])
		}
	}
	return res
}

func charToScore(chr string) int {
	var res int
	if unicode.IsUpper(rune(chr[0])) {
		res = int(chr[0]) - int('A') + 27
	} else {
		res = int(chr[0]) - int('a') + 1
	}
	return res
}

func part1() int {
	dat, _ := os.ReadFile("./day3-input")
	s := strings.Split(string(dat), "\n")
	var p1, p2 string
	var score int
	for i := 0; i < len(s)-1; i++ {
		p1 = s[i][len(s[i])/2:]
		p2 = s[i][:len(s[i])/2]
		score += charToScore(findSameItem(p1, p2))
	}
	return score
}

func part2() int {
	dat, _ := os.ReadFile("./day3-input")
	s := strings.Split(string(dat), "\n")
	var p1, p2, p3 string
	var score int
	for i := 0; i < (len(s)-1)/3; i++ {
		p1 = s[3*i]
		p2 = s[3*i+1]
		p3 = s[3*i+2]
		score += charToScore(findSameItemTriple(p1, p2, p3))
	}
	return score
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
	fmt.Println("DONE")
}
