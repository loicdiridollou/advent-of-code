package main

import (
	"strings"
)

func prepareData(input string) map[string][]string {
	caves := make(map[string][]string, 0)

	for _, path := range strings.Split(input, "\n") {
		if len(path) == 0 {
			continue
		}
		components := strings.Split(path, "-")
		if _, ok := caves[components[0]]; !ok {
			caves[components[0]] = make([]string, 0)
		}
		if components[1] != "start" {
			caves[components[0]] = append(caves[components[0]], components[1])
		}
		if _, ok := caves[components[1]]; !ok {
			caves[components[1]] = make([]string, 0)
		}
		if components[0] != "start" {
			caves[components[1]] = append(caves[components[1]], components[0])
		}
	}

	caves["end"] = make([]string, 0)
	return caves
}

func pathContains(path []string, cave string) bool {
	for _, past_cave := range path {
		if cave == past_cave {
			return true
		}
	}
	return false
}

func visitCaves(position string, caves map[string][]string, current_path []string) [][]string {
	list_paths := make([][]string, 0)

	for _, cave := range caves[position] {
		if cave == "end" {
			list_paths = append(list_paths, append(current_path, "end"))
		} else if canVisit(current_path, cave, 1) {
			list_paths = append(list_paths, visitCaves(cave, caves, append(current_path, cave))...)
		}
	}

	return list_paths
}

func part1(input string) int {
	caves := prepareData(input)
	paths := visitCaves("start", caves, []string{"start"})
	return len(paths)
}
