package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4")
  var dat []byte
  if len(os.Args) > 1 && os.Args[1] == "test" {
    dat, _ = os.ReadFile("./day4-test-input")
  } else {
    dat, _ = os.ReadFile("./day4-input")
  }
	fmt.Println("Part 1 result:", part1(string(dat)))
  fmt.Println("Part 2 result:", part2(string(dat)))
}
