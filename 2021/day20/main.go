package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("2021 - Day 20")
	dat, _ := os.ReadFile("./day20.input")

	fmt.Println("Part 1 result:", part1(string(dat)))
	fmt.Println("Part 2 result:", part2(string(dat)))
}
