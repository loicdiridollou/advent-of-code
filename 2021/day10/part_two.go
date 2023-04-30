package main

import (
	"sort"
)

func part2(input string) int {
	lines := prepareData(input)

	chr_map := make(map[string]string, 0)
	chr_map[")"] = "("
	chr_map["]"] = "["
	chr_map["}"] = "{"
	chr_map[">"] = "<"

	scores := make(map[string]int, 0)
	scores["("] = 1
	scores["["] = 2
	scores["{"] = 3
	scores["<"] = 4

	valid_lines := make([][]string, 0)

	for i := 0; i < len(lines); i++ {
		tmp_val := ""
		flag := true
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "(" || lines[i][j] == "[" || lines[i][j] == "{" || lines[i][j] == "<" {
				tmp_val += lines[i][j]
			} else if lines[i][j] == ")" || lines[i][j] == "]" || lines[i][j] == "}" || lines[i][j] == ">" {
				if chr_map[lines[i][j]] != string(tmp_val[len(tmp_val)-1]) {
					flag = false
					break
				}
				tmp_val = tmp_val[:len(tmp_val)-1]
			}
		}
		if flag {
			valid_lines = append(valid_lines, lines[i])
		}
	}

	to_complete := make([]string, 0)

	for i := 0; i < len(valid_lines); i++ {
		tmp_val := ""
		for j := 0; j < len(valid_lines[i]); j++ {
			if valid_lines[i][j] == "(" || valid_lines[i][j] == "[" || valid_lines[i][j] == "{" || valid_lines[i][j] == "<" {
				tmp_val += valid_lines[i][j]
			} else if valid_lines[i][j] == ")" || valid_lines[i][j] == "]" || valid_lines[i][j] == "}" || valid_lines[i][j] == ">" {
				tmp_val = tmp_val[:len(tmp_val)-1]
			}
		}
		to_complete = append(to_complete, tmp_val)
	}

	list_scores := make([]int, 0)
	for _, slc := range to_complete {
		value := 0
		for i := len(slc) - 1; i >= 0; i-- {
			value = 5*value + scores[string(slc[i])]
		}
		list_scores = append(list_scores, value)
	}

	sort.Slice(list_scores, func(i, j int) bool {
		return list_scores[i] < list_scores[j]
	})

	return list_scores[len(list_scores)/2]
}
