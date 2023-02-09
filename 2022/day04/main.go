package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4")
  dat, _ := os.ReadFile("./day04.input")
	fmt.Println("Part 1 result:", part1(string(dat)))
  fmt.Println("Part 2 result:", part2(string(dat)))
}
