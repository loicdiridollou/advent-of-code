package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

type Bliz struct {
	r, c int
	dir  string
}

type Exp struct {
	r, c int
}

func parseInput(input []string) ([]Bliz, [2]int, [2]int) {
	blizs := []Bliz{}
	dir := ""

	for r, el := range input {
		for c, chr := range el {
			if string(chr) != "." && string(chr) != "#" {
				switch string(chr) {
				case ">":
					dir = "E"
				case "<":
					dir = "W"
				case "^":
					dir = "N"
				case "v":
					dir = "S"
				}
				blizs = append(blizs, Bliz{r, c, dir})
			}
		}
	}

	return blizs, [2]int{0, 1}, [2]int{len(input) - 1, len(input[0]) - 2}
}

func findDim(input []string) (int, int) {
	return len(input) - 1, len(input[0]) - 1
}

func (bliz *Bliz) Move(limitb, limitr int) Bliz {
	dirs := map[string][2]int{"N": {-1, 0}, "S": {1, 0}, "W": {0, -1}, "E": {0, 1}}
	nr, nc := bliz.r+dirs[bliz.dir][0], bliz.c+dirs[bliz.dir][1]

	if nr == 0 {
		nr = limitb - 1
	} else if nr == limitb {
		nr = 1
	} else if nc == 0 {
		nc = limitr - 1
	} else if nc == limitr {
		nc = 1
	}

	return Bliz{nr, nc, bliz.dir}
}

func moveBliz(blizs []Bliz, limitb int, limitr int) []Bliz {
	res := []Bliz{}
	for _, blz := range blizs {
		res = append(res, blz.Move(limitb, limitr))
	}
	return res
}

func hashPositionExp(bliz Exp) string {
	return fmt.Sprint(bliz.r) + "_" + fmt.Sprint(bliz.c)
}

func hashPosition(bliz Bliz) string {
	return fmt.Sprint(bliz.r) + "_" + fmt.Sprint(bliz.c)
}

func hashBlizs(blizs []Bliz) map[string]bool {
	res := map[string]bool{}
	for _, bliz := range blizs {
		res[hashPosition(bliz)] = true
	}
	return res
}

func plotBlizs(blizs []Bliz, limitb, limitr int, exps []Exp) [][]string {
	res := make([][]string, 0)
	tmp := []string{}
	dirs := map[string]string{"N": "^", "S": "v", "W": "<", "E": ">"}
	for j := 0; j <= limitr; j++ {
		if j == 1 {
			tmp = append(tmp, ".")
		} else {
			tmp = append(tmp, "#")
		}
	}
	res = append(res, tmp)

	for i := 1; i < limitb; i++ {
		tmp := []string{"#"}

		for j := 1; j < limitr; j++ {
			tmp = append(tmp, ".")
		}
		tmp = append(tmp, "#")
		res = append(res, tmp)
	}
	tmp = []string{}
	for j := 0; j <= limitr; j++ {
		if j == limitr-1 {
			tmp = append(tmp, ".")
		} else {
			tmp = append(tmp, "#")
		}
	}

	res = append(res, tmp)

	for _, bliz := range blizs {
		res[bliz.r][bliz.c] = dirs[bliz.dir]
	}
	for _, exp := range exps {
		res[exp.r][exp.c] = "E"
	}
	return res
}

func moveExp(exp Exp) []Exp {
	dirs := map[string][2]int{"N": {-1, 0}, "S": {1, 0}, "W": {0, -1}, "E": {0, 1}}
	res := []Exp{exp}

	for _, coord := range dirs {
		nr, nc := exp.r+coord[0], exp.c+coord[1]
		res = append(res, Exp{nr, nc})
	}
	return res
}

func validateMoves(pot_moves []Exp, blizs_map map[string]bool, start [2]int, end [2]int, limitb, limitr int) []Exp {
	res := []Exp{}

	for _, exp := range pot_moves {
		if blizs_map[hashPositionExp(exp)] {
			continue
		} else if hashPositionExp(exp) == hashPositionExp(Exp{start[0], start[1]}) {
			res = append(res, exp)
		} else if hashPositionExp(exp) == hashPositionExp(Exp{end[0], end[1]}) {
			res = append(res, exp)
		} else if exp.r <= 0 || exp.r == limitb || exp.c <= 0 || exp.c == limitr {
			continue
		} else {
			res = append(res, exp)
		}
	}

	return res
}

func configBlizs(blizs []Bliz, round int, limitb, limitr int) map[int][]Bliz {
	res := map[int][]Bliz{0: blizs}

	for t := 1; t < round; t++ {
		res[t] = moveBliz(res[t-1], limitb, limitr)
	}

	return res
}

func hashMove(exp Exp, round int) string {
	return fmt.Sprint(exp.r) + "_" + fmt.Sprint(exp.c) + "_" + fmt.Sprint(round)
}

func part1() int {
	dat, _ := os.ReadFile("./day24-input")
	input := strings.Split(string(dat), "\n")
	input = input[:len(input)-1]

	blizs, start, end := parseInput(input)
	exp := Exp{start[0], start[1]}
	limitb, limitr := findDim(input)
	round := 0
	queue := []Exp{exp}
	arrived := false
	blizs_config := configBlizs(blizs, 300, limitb, limitr)
	queue_map := map[string]bool{}

	for round = 0; round < 300; round++ {
		blizs = blizs_config[round+1]
		blizs_map := hashBlizs(blizs)
		tmp_queue := []Exp{}
		for _, exp := range queue {
			if hashPositionExp(exp) == hashPositionExp(Exp{end[0], end[1]}) {
				arrived = true
				break
			}
			pot_moves := moveExp(exp)
			pot_moves = validateMoves(pot_moves, blizs_map, start, end, limitb, limitr)
			for _, mv := range pot_moves {
				if queue_map[hashMove(mv, round)] {
					continue
				} else {
					queue_map[hashMove(mv, round)] = true
					tmp_queue = append(tmp_queue, mv)
				}
			}
		}
		if arrived {
			break
		}
		queue = tmp_queue
	}

	return round
}
