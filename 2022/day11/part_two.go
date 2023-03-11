package main

import (
	"strings"
)

func part2(input string) int {
	s := strings.Split(input, "\n\n")
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
