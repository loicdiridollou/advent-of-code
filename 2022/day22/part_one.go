package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

)

func MaxInt(ls ...int) int {
  max := ls[0]

  for _, el := range ls {
    if el > max {
      max = el
    }
  }
  
  return max
}

func findStart(grid [][]string) [2]int {
  for idx_i, i := range grid {
    for idx_j, j := range i {
      if j == "." {
        return [2]int{idx_i, idx_j}
      }
    }
  }
  return [2]int{-1, -1}
}

func splitPath(path string) []string {
  res := make([]string, 0)
  tmp := ""
  for i := 0; i < len(path); i++ {
    if path[i] <= 'Z' && 'A' <= path[i] {
      res = append(res, tmp)
      tmp = ""
      res = append(res, string(path[i]))
    } else {
      tmp = tmp + string(path[i])
    }
  }
  return res
}
        

func part1() int {
	dat, _ := os.ReadFile("./day22-test-input")
	input := strings.Split(string(dat), "\n\n")
  maze := input[0]
  path := input[1]

  var max int

  for _, el := range strings.Split(string(maze), "\n") {
    if len(el) == 0 {
      continue
    }
    max = MaxInt(max, len(el))
  }
  
  var grid [][]string

  for _, el := range strings.Split(string(maze), "\n") {
    for len(el) < max {
      el = el + " "
    }
    grid = append(grid, strings.Split(el, ""))
  }


  start_pos := findStart(grid)

  fmt.Println(splitPath(path))
	return len(start_pos)
}
