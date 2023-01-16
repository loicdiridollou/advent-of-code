package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type transformer func(int) int

type monkey struct {
	num        int
	items      []int
	oper       func(old int) int
	test       int
	rule_true  int
	rule_false int
	done       int
}

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func mod(a int, t transformer) transformer {
	return func(b int) int {
		return t(b) % a
	}
}

func parseOper(oper string) func(old int) int {
	var fnc func(old int) int
	if strings.Contains(oper, "+") {
		sides := strings.Split(oper, " + ")
		if sides[0] == "old" && sides[1] == "old" {
			fnc = func(old int) int { return old + old }
		} else if sides[0] == "old" {
			fnc = func(old int) int { return old + convInt(sides[1]) }
		} else if sides[1] == "old" {
			fnc = func(old int) int { return convInt(sides[0]) + old }
		}
	} else if strings.Contains(oper, "*") {
		sides := strings.Split(oper, " * ")
		if sides[0] == "old" && sides[1] == "old" {
			fnc = func(old int) int { return old * old }
		} else if sides[0] == "old" {
			fnc = func(old int) int { return old * convInt(sides[1]) }
		} else if sides[1] == "old" {
			fnc = func(old int) int { return convInt(sides[0]) * old }
		}
	} else if strings.Contains(oper, "/") {
		sides := strings.Split(oper, " / ")
		if sides[0] == "old" && sides[1] == "old" {
			fnc = func(old int) int { return 1 }
		} else if sides[0] == "old" {
			fnc = func(old int) int { return old / convInt(sides[1]) }
		} else if sides[1] == "old" {
			fnc = func(old int) int { return convInt(sides[0]) / old }
		}
	}
	return fnc
}

func Play(num int, monkeys map[int]monkey, div int, inspections map[int]int) map[int]monkey {
	m := monkeys[num]
	for _, el := range m.items {
		inspections[num]++
		new_level := m.oper(el) / div
		if new_level%m.test == 0 {
			to_give := monkeys[m.rule_true]
			to_give.items = append(to_give.items, new_level)
			monkeys[m.rule_true] = to_give
		} else {
			to_give := monkeys[m.rule_false]
			to_give.items = append(to_give.items, new_level)
			monkeys[m.rule_false] = to_give
		}
	}
	m.items = m.items[:0]
	monkeys[num] = m
	return monkeys
}

func (m *monkey) step(monkey []*monkey) {
	for i := 0; i < len(m.items); i++ {
		new := m.oper(m.items[i])
		if new%m.test == 0 {
			monkey[m.rule_true].items = append(monkey[m.rule_true].items, new)
		} else {
			monkey[m.rule_false].items = append(monkey[m.rule_false].items, new)
		}
		m.done++
	}
	m.items = m.items[:0]
}

func topTwo(dic map[int]int) []int {
	lst := make([]int, 0)
	for _, el := range dic {
		lst = append(lst, el)
	}
	sort.Slice(lst, func(i, j int) bool {
		return lst[i] < lst[j]
	})
	return lst[len(lst)-2:]
}

func part1() int {
	dat, _ := os.ReadFile("./day11-input")
	s := strings.Split(string(dat), "\n\n")
	monkeys := make(map[int]monkey, 0)
	monkeys_list := make([]int, 0)
	inspections := make(map[int]int, 0)
	for _, el := range s {
		ss := strings.Split(string(el), "\n")
		num := string(strings.Split(string(ss[0]), " ")[1])
		num = string(num[:len(num)-1])
		items := string(strings.Split(ss[1], ": ")[1])
		itemsLst := make([]int, 0)
		for _, el := range strings.Split(items, ", ") {
			itemsLst = append(itemsLst, convInt(el))
		}
		oper := string(strings.Split(ss[2], ": new = ")[1])
		fnc := parseOper(oper)
		rule_div := string(strings.Split(ss[3], "by ")[1])
		rule_true := string(strings.Split(ss[4], "to monkey ")[1])
		rule_false := string(strings.Split(ss[5], "to monkey ")[1])
		m := monkey{
			num: convInt(num), items: itemsLst, oper: fnc, test: convInt(rule_div),
			rule_true: convInt(rule_true), rule_false: convInt(rule_false),
		}
		monkeys[convInt(num)] = m
		monkeys_list = append(monkeys_list, convInt(num))
		inspections[convInt(num)] = 0
	}
	for i := 0; i < 20; i++ {
		for _, monkey_num := range monkeys_list {
			monkeys = Play(monkey_num, monkeys, 3, inspections)
		}
	}
	tops := topTwo(inspections)
	return tops[0] * tops[1]
}

func part2() int {
	dat, _ := os.ReadFile("./day11-input")
	s := strings.Split(string(dat), "\n\n")
	monkeys := make(map[int]monkey, 0)
	monkeys_list := make([]int, 0)
	inspections := make(map[int]int, 0)
	for _, el := range s {
		ss := strings.Split(string(el), "\n")
		num := string(strings.Split(string(ss[0]), " ")[1])
		num = string(num[:len(num)-1])
		items := string(strings.Split(ss[1], ": ")[1])
		itemsLst := make([]int, 0)
		for _, el := range strings.Split(items, ", ") {
			itemsLst = append(itemsLst, convInt(el))
		}
		oper := string(strings.Split(ss[2], ": new = ")[1])
		fnc := parseOper(oper)
		rule_div := string(strings.Split(ss[3], "by ")[1])
		p := 2 * 17 * 19 * 3 * 5 * 13 * 7 * 11
		fnc = mod(p, fnc)
		rule_true := string(strings.Split(ss[4], "to monkey ")[1])
		rule_false := string(strings.Split(ss[5], "to monkey ")[1])
		m := monkey{
			num: convInt(num), items: itemsLst, oper: fnc, test: convInt(rule_div),
			rule_true: convInt(rule_true), rule_false: convInt(rule_false), done: 0,
		}
		monkeys_list = append(monkeys_list, convInt(num))
		monkeys[convInt(num)] = m
		inspections[convInt(num)] = 0
	}
	for i := 0; i < 10000; i++ {
		for _, monkey_num := range monkeys_list {
			monkeys = Play(monkey_num, monkeys, 1, inspections)
		}
	}
	tops := topTwo(inspections)
	return tops[0] * tops[1]
}

func main() {
	fmt.Println("Part 1 result:", part1())
	fmt.Println("Part 2 result:", part2())
	fmt.Println("DONE")
}
