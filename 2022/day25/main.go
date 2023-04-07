package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 25")
	var dat []byte
	dat, _ = os.ReadFile("./day25.input")

	fmt.Println("Part 1 result:", part1(string(dat)))
}
