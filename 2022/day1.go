package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addArray(numbs []string) int {
	result := 0
	for _, numb := range numbs {
		num, _ := strconv.Atoi(numb)
		result += num
	}
	return result
}

func addArrayInt(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}

func maxArray(numbs []int) int {
	for j := 1; j < len(numbs); j++ {
		if numbs[j] > numbs[0] {
			numbs[0] = numbs[j]
		}
	}
	return numbs[0]
}

// part1 function
func part1() int {
	dat, _ := os.ReadFile("./day1-input")
	s := strings.Split(string(dat), "\n\n")
	var res []int

	for i := 0; i < len(s); i++ {
		res = append(res, addArray(strings.Split(s[i], "\n")))
	}
	return maxArray(res)
}

// part2 function
func part2() int {
	dat, _ := os.ReadFile("./day1-input")
	s := strings.Split(string(dat), "\n\n")
	var res []int

	for i := 0; i < len(s); i++ {
		res = append(res, addArray(strings.Split(s[i], "\n")))
	}
	sort.Ints(res)
	return addArrayInt(res[len(res)-3:])
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
	fmt.Println("DONE")
}
