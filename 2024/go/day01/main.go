package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 1")
	var dat []byte
	dat, _ = os.ReadFile("./input.txt")
	fmt.Println("Part 1 result:", part1(string(dat)))
	fmt.Println("Part 2 result:", part2(string(dat)))
}
