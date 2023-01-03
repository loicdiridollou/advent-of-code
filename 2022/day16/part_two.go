package main

import (
  "os"
  "sort"
  "strings"
)


func part2() int {
	dat, _ := os.ReadFile("./day16-input")
	s := strings.Split(string(dat), "\n")

	var readings []Valve
  for _, ln := range s[:len(s)-1] {
		readings = append(readings, parseLine(ln))
	}

  paths := DFS(readings, 26)


  sorted := make([]BitSet, 0, len(paths))
	for key := range paths {
		sorted = append(sorted, key)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return paths[sorted[i]] > paths[sorted[j]]
	})

  max := 0
  for _, open1 := range sorted {
    pressure1 := paths[open1]
    for _, open2 := range sorted {
      if open1 & open2 == 0 {
        if sum := pressure1 + paths[open2]; max < sum {
          max = sum
        } else if sum < max {
          break
        }
      }
    }
  }

  return max
}
