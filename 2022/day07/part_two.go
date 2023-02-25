package main

import "strings"

func part2(input string) int {
	s := strings.Split(input, "\n")
	dir, dir_lst := createMap(s)

	finalSize := 0
	dirSize := make(map[string]int, 0)
	for _, el := range dir_lst {
		tmpSize := computeSize(el, dir)
		dirSize[el] = tmpSize
	}

	reqSize := 30000000 - 70000000 + dirSize[".//"]
	finalSize = 70000000
	for _, el := range dir_lst {
		tmpSize := computeSize(el, dir)
		if tmpSize > reqSize && tmpSize < finalSize {
			finalSize = tmpSize
		}
	}

	return finalSize
}
