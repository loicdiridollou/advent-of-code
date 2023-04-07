package main

import (
	"strings"
)

func part2(input string) int {
	input_data := strings.Split(input, "\n")

	elves := parseInput(input_data)
	rot := 0
	score := 0
	round := 0

	for round = 0; round < 1000; round++ {
		moves, list_move := findMoves(elves, rot)

		score = 0
		for i, elf := range elves {
			move, moving := moves[hashPosition(elf)]
			if !moving {
				continue
			} else if list_move[move] != 1 {
				continue
			} else {
				score++
				elves[i] = elf.Move(parseMove(move))
			}
		}
		if score == 0 {
			break
		}
		rot = (rot + 1) % 4
	}

	return round + 1
}
