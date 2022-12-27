package main

import (
	"fmt"
	"os"
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

func part1() int {
	dat, _ := os.ReadFile("./day7-input")
	s := strings.Split(string(dat), "\n")
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

func part2() int {
	dat, _ := os.ReadFile("./day7-input")
	s := strings.Split(string(dat), "\n")
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

func main() {
	fmt.Println("Part 1 result:", part1())
	fmt.Println("Part 2 result:", part2())
	fmt.Println("DONE")
}
