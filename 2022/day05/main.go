package main

import (
	"fmt"
	"os"
  "gitlab.com/loicdiridollou/advent-of-code/util/mathemat"
)

func main() {
	fmt.Println("Day 5")
	var dat []byte
	if len(os.Args) > 1 && os.Args[1] == "test" {
		dat, _ = os.ReadFile("./day5-test-input")
	} else {
		dat, _ = os.ReadFile("./day5-input")
	}
  mathemat.Add(1, 2)

  fmt.Println(part1(string(dat)))
  fmt.Println(part2(string(dat)))
}
