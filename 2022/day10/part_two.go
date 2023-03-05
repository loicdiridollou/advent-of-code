package main

import (
	"strings"
)

func part2(input string) [6][]string {
	s := strings.Split(input, "\n")
	screen := [6][]string{}
	strength := 1
	for i := range screen {
		d := make([]string, 0)
		for j := 0; j < 40; j++ {
			d = append(d, ".")
		}
		screen[i] = d
	}
	cycle := 1
	idx := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != "noop" {
			if strength-1 <= idx%40 && idx%40 <= strength+1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
			if strength-1 <= idx%40 && idx%40 <= strength+1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
			toAdd := convInt(strings.Split(s[i], " ")[1])
			strength = strength + toAdd
		} else {
			if strength-1 <= idx%40 && idx%40 <= strength+1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
		}
	}
	return screen
}
