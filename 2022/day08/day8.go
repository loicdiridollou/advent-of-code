package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func MaxArray(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func formatInput(s []string) [][]int {
	tbl := make([][]int, len(s)-1)
	for i := range tbl {
		for j := range string(s[i]) {
			tbl[i] = append(tbl[i], convInt(string(s[i][j])))
		}
	}
	return tbl
}

func extColumn(board [][]int, columnIndex int) []int {
	column := make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return column
}

func isVisible(i int, j int, tbl [][]int) bool {
	col := extColumn(tbl, j)
	var res bool
	if MaxArray(tbl[i][:j]) < tbl[i][j] {
		res = true
	} else if MaxArray(tbl[i][j+1:]) < tbl[i][j] {
		res = true
	} else if MaxArray(col[:i]) < tbl[i][j] {
		res = true
	} else if MaxArray(col[i+1:]) < tbl[i][j] {
		res = true
	} else {
		res = false
	}
	return res
}

func Reverse(input []int) []int {
	inputLen := len(input)
	output := make([]int, inputLen)
	for i, n := range input {
		j := inputLen - i - 1
		output[j] = n
	}
	return output
}

func computeScore(i int, j int, tbl [][]int) int {
	scores := make([]int, 0)
	col := extColumn(tbl, j)
	t1 := tbl[i][:j]
	t2 := tbl[i][j+1:]
	t3 := col[:i]
	t4 := col[i+1:]
	tot := [4][]int{Reverse(t1), t2, Reverse(t3), t4}
	for _, lt := range tot {
		flag := false
		score := 0
		for _, k := range lt {
			if !flag {
				if k < tbl[i][j] {
					score++
				} else if k == tbl[i][j] {
					score++
					flag = true
				} else {
					score++
					flag = true
				}
			}
		}
		scores = append(scores, score)
	}
	return scores[0] * scores[1] * scores[2] * scores[3]
}

func part1() int {
	dat, _ := os.ReadFile("./day8-input")
	s := strings.Split(string(dat), "\n")
	tbl := formatInput(s)
	totVisible := 2*len(tbl) + 2*len(tbl[0]) - 4

	for i := 1; i < len(tbl)-1; i++ {
		for j := 1; j < len(tbl[0])-1; j++ {
			if isVisible(i, j, tbl) {
				totVisible++
			}
		}
	}
	return totVisible
}

func part2() int {
	dat, _ := os.ReadFile("./day8-input")
	s := strings.Split(string(dat), "\n")
	tbl := formatInput(s)
	totVisible := make([]int, 0)

	for i := 1; i < len(tbl)-1; i++ {
		for j := 1; j < len(tbl[0])-1; j++ {
			totVisible = append(totVisible, computeScore(i, j, tbl))
		}
	}
	return MaxArray(totVisible)
}

func main() {
	fmt.Println("Part 1 result:", part1())
	fmt.Println("Part 2 result:", part2())
	fmt.Println("DONE")
}
