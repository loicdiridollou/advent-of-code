package main

import (
	"strings"
)

func part2(input string) int {
	s := strings.Split(input, "\n\n")

	numbers := s[0]
	boards := prepareData(s[1:])
	num := 0
	for _, number := range strings.Split(numbers, ",") {
		for _, board := range boards {
			if win, _ := board.HasWon(); win {
				continue
			}
			board.Play(string(number))
			win, val := board.HasWon()
			if win && num == len(boards)-1 {
				return StrToInt(number) * val
			} else if win {
				num++
			}
		}
	}
	return 0
}
