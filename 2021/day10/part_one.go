package main

import (
	"strings"
)

func prepareData(input string) [][]string {
	res := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		tmp_res := make([]string, 0)
		for _, chr := range line {
			tmp_res = append(tmp_res, string(chr))
		}
		res = append(res, tmp_res)
	}
	return res
}

func part1(input string) int {
	lines := prepareData(input)
	res := make([]string, 0)

	chr_map := make(map[string]string, 0)
	chr_map[")"] = "("
	chr_map["]"] = "["
	chr_map["}"] = "{"
	chr_map[">"] = "<"

	scores := make(map[string]int, 0)
	scores[")"] = 3
	scores["]"] = 57
	scores["}"] = 1197
	scores[">"] = 25137

	for i := 0; i < len(lines); i++ {
		tmp_val := ""
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "(" || lines[i][j] == "[" || lines[i][j] == "{" || lines[i][j] == "<" {
				tmp_val += lines[i][j]
			} else if lines[i][j] == ")" || lines[i][j] == "]" || lines[i][j] == "}" || lines[i][j] == ">" {
				if chr_map[lines[i][j]] != string(tmp_val[len(tmp_val)-1]) {
					res = append(res, lines[i][j])
					break
				}
				tmp_val = tmp_val[:len(tmp_val)-1]
			}
		}
	}

	value := 0
	for _, el := range res {
		value += scores[el]
	}

	return value
}
