package main

import (
	"fmt"
	"strings"
)

func canVisit(path []string, cave string, num_visits int) bool {
	tmp_visit := 0
	if cave == "start" {
		return false
	}
	if cave == "end" {
		return false
	}
	for _, past_cave := range path {
		if strings.ToLower(past_cave) == past_cave && cave == past_cave {
			fmt.Println(past_cave)
			tmp_visit++
		}
	}
	return tmp_visit < num_visits
}

func visitCavesDouble(position string, caves map[string][]string, current_path []string) [][]string {
	list_paths := make([][]string, 0)

	for _, cave := range caves[position] {
		if cave == "end" && current_path[len(current_path)-1] != "end" {
			list_paths = append(list_paths, append(current_path, "end"))
		} else if canVisit(current_path, cave, 2) {
			new_path := append(current_path, cave)
			new_paths := visitCavesDouble(cave, caves, new_path)
			if len(new_paths) > 0 {
				fmt.Println(current_path, new_paths)
				list_paths = append(list_paths, new_paths...)
			}
		}
	}

	return list_paths
}

func part2(input string) int {
	caves := prepareData(input)
	paths := visitCavesDouble("start", caves, []string{"start"})
	return len(paths)
}
