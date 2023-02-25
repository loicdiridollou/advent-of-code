package main

import (
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func computeSize(path string, dir map[string][]string) int {
	lst := dir[path]
	totSize := 0
	for _, el := range lst {
		p1 := strings.Split(el, " ")[1]
		p2 := strings.Split(el, " ")[0]
		if string(p2) == "dir" {
			totSize = totSize + computeSize(path+"/"+p1, dir)
		} else {
			totSize += convInt(p2)
		}
	}
	return totSize
}

func createMap(s []string) (map[string][]string, []string) {
	dir := make(map[string][]string, 0)
	dir_lst := make([]string, 0)
	path := "."
	var tmp []string
	for i := 0; i < len(s)-1; i++ {
		if string(s[i][0]) == "$" {
			if string(s[i]) == "$ cd .." {
				tmp = strings.Split(path, "/")
				path = strings.Join(tmp[:len(tmp)-1], "/")
			} else if string(s[i][2:4]) == "cd" {
				path = path + "/" + string(s[i][5:])
				dir_lst = append(dir_lst, path)
			}
		} else {
			_, exists := dir[path]
			if !exists {
				dir[path] = make([]string, 0)
			}
			dir[path] = append(dir[path], s[i])
		}
	}

	return dir, dir_lst
}

func part1(input string) int {
	s := strings.Split(input, "\n")
	dir, dir_lst := createMap(s)

	finalSize := 0
	for _, el := range dir_lst {
		tmpSize := computeSize(el, dir)
		if tmpSize < 100000 {
			finalSize += tmpSize
		}
	}
	return finalSize
}
