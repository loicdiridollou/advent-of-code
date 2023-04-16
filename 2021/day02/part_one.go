package main

import (
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

type Postion struct {
	x, y int
}

func part1(input string) int {
	s := strings.Split(input, "\n")
	pos := Postion{0, 0}

	for _, el := range s {
		if len(el) == 0 {
			continue
		}
		a := strings.Split(el, " ")
		switch a[0] {
		case "forward":
			pos.x += StrToInt(a[1])
		case "up":
			pos.y -= StrToInt(a[1])
		case "down":
			pos.y += StrToInt(a[1])
		}
	}

	return pos.x * pos.y
}
