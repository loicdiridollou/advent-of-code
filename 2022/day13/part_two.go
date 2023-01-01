package main

import (
	"os"
	"sort"
	"strings"
)

func part2() int {
  dat, _ := os.ReadFile("./day13-input")
	s := strings.Split(strings.Join(strings.Split(string(dat), "\n\n"), "\n"), "\n")
  s = s[:len(s)-1]
  var packages []tree
  for _, st := range s {
    packages = append(packages, readTree(st))
  }
  packages = append(packages, readTree("[[2]]"))
  packages = append(packages, readTree("[[6]]"))
  sort.Slice(packages, func(i, j int) bool {return checkOrder(packages[i], packages[j]) == 1})
  decoder := 1
  for i, p := range packages {
    if checkOrder(p, readTree("[[2]]")) == 0 || checkOrder(p, readTree("[[6]]")) == 0 {
      decoder *= i+1
    }
  }
	return decoder
}
