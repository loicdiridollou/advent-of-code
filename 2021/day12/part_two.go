package main

import (
	"strings"
)

func sliceContains(slice []string, val string) int {
	num := 0
	if strings.ToUpper(val) == val {
		return 0
	}
	for _, el := range slice {
		if el == val {
			num++
		}
	}
	return num
}

func Counter(s []string) map[string]int {
	dic := make(map[string]int)

	for _, el := range s {
		_, exists := dic[string(el)]
		if !exists {
			dic[string(el)] = 0
		}
		dic[string(el)] += 1
	}

	return dic
}

func smallCaveVisitedTwice(path []string) bool {
	counter := Counter(path)
	for j, k := range counter {
		if k == 2 && strings.ToLower(j) == j {
			return false
		}
	}
	return true
}

func visitCavesDouble(position string, caves map[string][]string, current_path []string) [][]string {
	list_paths := make([][]string, 0)

	for _, cave := range caves[position] {
		if cave == "end" {
			list_paths = append(list_paths, append(current_path, "end"))
		} else if sliceContains(current_path, cave) < 2 && smallCaveVisitedTwice(current_path) {
			list_paths = append(list_paths, visitCavesDouble(cave, caves, append(current_path, cave))...)
		} else if sliceContains(current_path, cave) == 0 {
			list_paths = append(list_paths, visitCavesDouble(cave, caves, append(current_path, cave))...)
		}
	}

	return list_paths
}

func part2(input string) int {
	caves := prepareData(input)
	paths := visitCavesDouble("start", caves, []string{"start"})
	return len(paths)
}
