package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 15")
	var dat []byte
	dat, _ = os.ReadFile("./day15.input")

	fmt.Println("Part 1 result:", part1(string(dat), 2000000))
	fmt.Println("Part 2 result:", part2(string(dat), 4000000))
}
