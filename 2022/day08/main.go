package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 8")
	var dat []byte
	dat, _ = os.ReadFile("./day08.input")

  fmt.Println(part1(string(dat)))
  fmt.Println(part2(string(dat)))
}
