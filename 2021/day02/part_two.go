package main

import (
	"strings"
)

type PosAim struct {
	x, y, aim int
}

func part2(input string) int {
	s := strings.Split(input, "\n")
	pos := PosAim{0, 0, 0}

	for _, el := range s {
		if len(el) == 0 {
			continue
		}
		a := strings.Split(el, " ")
		switch a[0] {
		case "forward":
			pos.x += StrToInt(a[1])
			pos.y += StrToInt(a[1]) * pos.aim
		case "up":
			pos.aim -= StrToInt(a[1])
		case "down":
			pos.aim += StrToInt(a[1])
		}
	}

	return pos.x * pos.y
}
