package main

import (
  "os"
  "strings"
)

func part2() int {
	dat, _ := os.ReadFile("./day2-input")
	s := strings.Split(string(dat), "\n")
	res := ""
	play := ""
	score := 0
	for i := 0; i < len(s)-1; i++ {
		play = adaptPlay(string(s[i][0]), string(s[i][2]))
		res = gameResult(string(s[i][0]), play)
		if res == "Win" {
			score += 6
		} else if res == "Tie" {
			score += 3
		}
		if play == "X" {
			score += 1
		} else if play == "Y" {
			score += 2
		} else {
			score += 3
		}
	}
	return score
}
