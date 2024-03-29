package main

import (
	"strings"
)

func part2(input string) int {
	input_data := strings.Split(input, "\n")
	input_data = input_data[:len(input_data)-1]

	blizs, start, end := parseInput(input_data)
	exp := Exp{start[0], start[1]}
	limitb, limitr := findDim(input_data)
	round := 0
	queue := []Exp{exp}
	arrived := false
	blizs_config := configBlizs(blizs, 1000, limitb, limitr)
	queue_map := map[string]bool{}

	for round = 0; round < 1000; round++ {
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

	queue = []Exp{{end[0], end[1]}}
	arrived = false
	queue_map = map[string]bool{}
	round2 := round
	for round = round2; round < 1000; round++ {
		blizs = blizs_config[round+1]
		blizs_map := hashBlizs(blizs)
		tmp_queue := []Exp{}
		for _, exp := range queue {
			if hashPositionExp(exp) == hashPositionExp(Exp{start[0], start[1]}) {
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

	queue = []Exp{{start[0], start[1]}}
	arrived = false
	queue_map = map[string]bool{}
	round2 = round

	for round = round2; round < 1000; round++ {
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
