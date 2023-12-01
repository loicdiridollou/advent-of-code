package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

type Operation struct {
	type_op string
	v1, v2  string
}

func prepareData(input string) ([]Operation, int) {
	res := make([]Operation, 0)
	r, _ := regexp.Compile(`(\w+) (\w+) (\w)?`)
	num_input := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		matches := r.FindStringSubmatch(line)
		if len(matches) == 4 {
			res = append(res, Operation{matches[1], matches[2], matches[3]})
		} else {
			res = append(res, Operation{matches[1], matches[2], ""})
			num_input++
		}
	}
	return res, num_input
}

func part1(input string) int {
	operations, num_input := prepareData(input)
	fmt.Println(operations, num_input)

	return 0
}
