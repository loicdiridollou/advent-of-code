package main

import (
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func prepareData(input string) []int {
	res := make([]int, 0)
	r, _ := regexp.Compile(`Player \d+ starting position: (\d+)`)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		res = append(res, StrToInt(matches[1]))
	}
	return res
}

func adjustRoll(roll int) int {
	if roll > 100 {
		return roll - 100
	}
	return roll
}

func rollDice(last_roll int) (int, int) {
	sum := adjustRoll(last_roll+1) + adjustRoll(last_roll+2) + adjustRoll(last_roll+3)
	return sum, adjustRoll(last_roll + 3)
}

func part1(input string) int {
	positions := prepareData(input)
	p1, p2 := positions[0], positions[1]
	s1, s2 := 0, 0
	last_roll := 0
	num_roll := 0
	var s_tmp int

	for (s1 < 1000) || (s2 < 1000) {
		s_tmp, last_roll = rollDice(last_roll)
		p1 += (s_tmp-1)%10 + 1
		p1 = (p1-1)%10 + 1
		s1 += p1
		if s1 >= 1000 {
			num_roll += 3
			break
		}
		s_tmp, last_roll = rollDice(last_roll)
		p2 += (s_tmp-1)%10 + 1
		p2 = (p2-1)%10 + 1
		s2 += p2

		num_roll += 6
	}

	if s1 < 1000 {
		return s1 * num_roll
	} else {
		return s2 * num_roll
	}
}
