package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Day 11")
	var dat []byte
	dat, _ = os.ReadFile("./input.txt")
	start := time.Now()
	fmt.Println("Part 1 result:", part1(string(dat)))
	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	start = time.Now()
	fmt.Println("Part 2 result:", part2(string(dat)))
	elapsed = time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
}
