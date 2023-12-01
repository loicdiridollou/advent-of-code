package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("2021 - Day 23")
	dat, _ := os.ReadFile("./day23.testinput")

	fmt.Println("Part 1 result:", part1(string(dat)))
	fmt.Println("Part 2 result:", part2(string(dat)))
}
