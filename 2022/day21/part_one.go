package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type Monkey struct {
	id   string
	num  int
	wait string
}

func part1(input string) int {
	input_data := strings.Split(input, "\n")
	input_data = input_data[:len(input_data)-1]
	var monkeys []Monkey
	monkeys_map := make(map[string]int, 0)

	for _, el := range input_data {
		p := strings.Split(string(el), ": ")
		num, err := strconv.Atoi(p[1])
		if err != nil {
			monkey := Monkey{p[0], 0, p[1]}
			monkeys = append(monkeys, monkey)
		} else {
			monkeys_map[p[0]] = num
		}
	}

	for len(monkeys) > 0 {
		monkey := monkeys[0]
		monkeys = monkeys[1:]
		switch string(monkey.wait[5]) {
		case "+":
			p := strings.Split(string(monkey.wait), " + ")
			m1, f1 := monkeys_map[p[0]]
			m2, f2 := monkeys_map[p[1]]
			if f1 && f2 {
				monkeys_map[monkey.id] = m1 + m2
			} else {
				monkeys = append(monkeys, monkey)
			}
		case "-":
			p := strings.Split(string(monkey.wait), " - ")
			m1, f1 := monkeys_map[p[0]]
			m2, f2 := monkeys_map[p[1]]
			if f1 && f2 {
				monkeys_map[monkey.id] = m1 - m2
			} else {
				monkeys = append(monkeys, monkey)
			}
		case "*":
			p := strings.Split(string(monkey.wait), " * ")
			m1, f1 := monkeys_map[p[0]]
			m2, f2 := monkeys_map[p[1]]
			if f1 && f2 {
				monkeys_map[monkey.id] = m1 * m2
			} else {
				monkeys = append(monkeys, monkey)
			}
		case "/":
			p := strings.Split(string(monkey.wait), " / ")
			m1, f1 := monkeys_map[p[0]]
			m2, f2 := monkeys_map[p[1]]
			if f1 && f2 {
				monkeys_map[monkey.id] = m1 / m2
			} else {
				monkeys = append(monkeys, monkey)
			}
		}
	}

	return monkeys_map["root"]
}
