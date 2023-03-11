package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 11")
	var dat []byte
	dat, _ = os.ReadFile("./day11.input")

	fmt.Println(part1(string(dat)))
	fmt.Println(part2(string(dat)))
}
