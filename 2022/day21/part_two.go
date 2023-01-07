package main

import (
	_ "embed"
	"os"
	"strconv"
	"strings"
)

func parseLine(input []string) map[string]Monkey {
	monkeys := make(map[string]Monkey, 0)

	for _, el := range input {
		p := strings.Split(string(el), ": ")
		num, err := strconv.Atoi(p[1])
		if err != nil {
			monkey := Monkey{p[0], 0, p[1]}
			monkeys[string(p[0])] = monkey
		} else {
			monkeys[p[0]] = Monkey{p[0], num, ""}
		}
	}
	return monkeys
}

func find(name string, monkeys map[string]Monkey, h float64) float64 {
	monkey := monkeys[name]
	var val float64

	if name == "humn" && h >= 0 {
		return float64(h)
	} else {
		if monkey.wait == "" {
			return float64(monkey.num)
		} else {
			m1 := find(string(monkey.wait[:4]), monkeys, h)
			m2 := find(string(monkey.wait[7:]), monkeys, h)
			if string(monkey.wait[5]) == "+" {
				val = m1 + m2
			} else if string(monkey.wait[5]) == "-" {
				val = m1 - m2
			} else if string(monkey.wait[5]) == "*" {
				val = m1 * m2
			} else if string(monkey.wait[5]) == "/" {
				val = m1 / m2
			}
		}
	}
	return val
}

func part2() float64 {
	dat, _ := os.ReadFile("./day21-input")
	input := strings.Split(string(dat), "\n")
	input = input[:len(input)-1]
	monkeys := parseLine(input)

	var lo, hi float64 = 0, 1e13
	target := find(string(monkeys["root"].wait[7:]), monkeys, 0)
	var mid float64

	for lo < hi {
		mid = (hi + lo) / 2
		f1 := find(string(monkeys["root"].wait[:4]), monkeys, mid)
		if target-f1 < 0 {
			lo = mid
		} else if f1 == target {
			break
		} else {
			hi = mid
		}
	}

	return mid
}
