package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 21")
	var dat []byte
	dat, _ = os.ReadFile("./day21.input")

	fmt.Println("Part 1 result:", part1(string(dat)))
	fmt.Println("Part 2 result:", part2(string(dat)))
}
