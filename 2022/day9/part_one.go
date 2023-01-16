package main

import (
	_ "embed"
	"fmt"
	"os"
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

func (coord *Coord) Adjust(head Coord) {
  fmt.Println(coord, head)
}

func part1() int {
	dat, _ := os.ReadFile("./day9-test-input")
	input := strings.Split(string(dat), "\n")
	moves := parseInput(input)
  head := Coord{0, 0}
  tail := Coord{0, 0}
  tail_positions := map[string]bool{}
	fmt.Println(moves, head, tail)
  fmt.Println(tail.Hash())



	return len(tail_positions)
}
