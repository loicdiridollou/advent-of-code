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

func part1() int {
	dat, _ := os.ReadFile("./day10-input")
	s := strings.Split(string(dat), "\n")
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

func part2() int {
	dat, _ := os.ReadFile("./day10-input")
	s := strings.Split(string(dat), "\n")
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
	for i := 0; i < len(s) - 1; i++ {
		if s[i] != "noop" {
			if strength - 1  <= idx%40 && idx%40 <= strength + 1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
			if strength - 1  <= idx%40 && idx%40 <= strength + 1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
			toAdd := convInt(strings.Split(s[i], " ")[1])
			strength = strength + toAdd
		} else {
			if strength - 1  <= idx%40 && idx%40 <= strength + 1 {
				screen[idx/40][idx%40] = "#"
			} else {
				screen[idx/40][idx%40] = "."
			}
			cycle++
			idx++
		}
	}
	fmt.Println(screen)
	return 0
}

func main() {
	fmt.Println("Part 1 result:", part1())
	fmt.Println("Part 2 result:", part2())
	fmt.Println("DONE")
}
