package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("2021 - Day 12")
	dat, _ := os.ReadFile("./day12.testinput")
	fmt.Println("Part 1 result:", part1(string(dat)))
	fmt.Println("Part 2 result:", part2(string(dat)))
}
