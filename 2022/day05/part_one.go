package main

import (
	"regexp"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func cleanSlice(slice [][]string) [][]string {
	var res [][]string
	for _, s := range slice {
		var r []string
		for _, str := range s {
			if str != " " {
				r = append(r, str)
			}
		}
		res = append(res, r)
	}
	return res
}

func processStack(p1 []string) [][]string {
	numCol := (len(p1[0]) + 1) / 4
	numLine := len(p1) - 1

	var sl [][]string
	for i := numLine - 1; i >= 0; i-- {
		var tmp []string
		for j := 0; j < numCol; j++ {
			el := string(p1[i][4*j+1])
			tmp = append(tmp, el)
		}
		sl = append(sl, tmp)
	}
	return cleanSlice(transpose(sl))
}

func convListToInt(slice []string) []int {
	var res []int
	for _, el := range slice {
		res = append(res, convInt(el))
	}
	return res
}

func processInstructions(p2 []string) [][]int {
	var res [][]int
	r, _ := regexp.Compile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	for i := 0; i < len(p2)-1; i++ {
		res = append(res, convListToInt(r.FindStringSubmatch(p2[i])[1:]))
	}
	return res
}

func Reverse(input []string) []string {
	inputLen := len(input)
	output := make([]string, inputLen)
	for i, n := range input {
		j := inputLen - i - 1
		output[j] = n
	}
	return output
}

func executeCommands(slice [][]string, cmds [][]int, reverse bool) [][]string {
	for _, cmd := range cmds {
		if reverse {
			slice[cmd[2]-1] = append(slice[cmd[2]-1], Reverse(slice[cmd[1]-1][len(slice[cmd[1]-1])-cmd[0]:])...)
		} else {
			slice[cmd[2]-1] = append(slice[cmd[2]-1], slice[cmd[1]-1][len(slice[cmd[1]-1])-cmd[0]:]...)
		}
		slice[cmd[1]-1] = slice[cmd[1]-1][:len(slice[cmd[1]-1])-cmd[0]]
	}
	return slice
}

func extractTop(slice [][]string) string {
	var res []string
	for _, s := range slice {
		res = append(res, s[len(s)-1])
	}
	return strings.Join(res, "")
}

func part1(input string) string {
	s := strings.Split(input, "\n\n")
	p1 := strings.Split(s[0], "\n")
	p2 := strings.Split(s[1], "\n")
	sl := processStack(p1)
	cmd := processInstructions(p2)
	res := executeCommands(sl, cmd, true)
	return extractTop(res)
}
