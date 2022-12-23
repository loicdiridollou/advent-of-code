package main

import (
	"fmt"
	"os"
)

func allUnique(dic map[string]int, num int) bool {
	score := 0
	for _, el := range dic {
		if el != 0 {
			score++
		}
	}
	return (score == num)
}

func part1() int {
	dat, _ := os.ReadFile("./day6-input")
	s := string(dat)
	dic := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		if i >= 4 {
			dic[string(s[i-4])] = dic[string(s[i-4])] - 1
		}
		dic[string(s[i])] = dic[string(s[i])] + 1
		if allUnique(dic, 4) {
      return i + 1
    }
	}
	return -1
}

func part2() int {
	dat, _ := os.ReadFile("./day6-input")
	s := string(dat)
	dic := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		if i >= 14 {
			dic[string(s[i-14])] = dic[string(s[i-14])] - 1
		}
		dic[string(s[i])] = dic[string(s[i])] + 1
		if allUnique(dic, 14) {
      return i + 1
    }
	}
	return -1
}

func main() {
	fmt.Println("Part 1 result:", part1())
	fmt.Println("Part 2 result:", part2())
	fmt.Println("DONE")
}
