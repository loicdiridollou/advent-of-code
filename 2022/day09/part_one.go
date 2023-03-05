package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Move struct {
	dir    string
	length int
}

func convInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

type Coord struct {
	r, c int
}

func parseInput(input []string) []Move {
	moves := []Move{}

	for _, move := range input {
		if len(move) == 0 {
			continue
		}
		spl := strings.Split(move, " ")
		moves = append(moves, Move{spl[0], convInt(spl[1])})
	}

	return moves
}

func (coord Coord) Hash() string {
	return fmt.Sprint(coord.r) + "_" + fmt.Sprint(coord.c)
}

func (coord *Coord) Move(dir string) {
	switch dir {
	case "U":
		coord.r = coord.r - 1
	case "D":
		coord.r = coord.r + 1
	case "L":
		coord.c = coord.c - 1
	case "R":
		coord.c = coord.c + 1
	}
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func (coord *Coord) Adjust(head Coord) {
	dr, dc := coord.r-head.r, coord.c-head.c
	if Abs(dr) <= 1 && Abs(dc) <= 1 {
		return
	} else if Abs(dr) == 2 && dc == 0 {
		if dr > 0 {
			coord.r--
		} else {
			coord.r++
		}
	} else if Abs(dc) == 2 && dr == 0 {
		if dc < 0 {
			coord.c++
		} else {
			coord.c--
		}
	} else {
		if dr > 0 && dc > 0 {
			coord.r--
			coord.c--
		} else if dr > 0 && dc < 0 {
			coord.r--
			coord.c++
		} else if dr < 0 && dc > 0 {
			coord.r++
			coord.c--
		} else if dr < 0 && dc < 0 {
			coord.r++
			coord.c++
		}
	}
}

func part1(input string) int {
	moves := parseInput(strings.Split(input, "\n"))
	head := Coord{0, 0}
	tail := Coord{0, 0}
	tail_positions := map[string]bool{tail.Hash(): true}

	for _, move := range moves {
		for ln := 0; ln < move.length; ln++ {
			head.Move(move.dir)
			tail.Adjust(head)
			tail_positions[tail.Hash()] = true
		}
	}

	return len(tail_positions)
}
