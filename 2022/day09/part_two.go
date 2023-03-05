package main

import "strings"

func part2(input string) int {
	moves := parseInput(strings.Split(input, "\n"))
	rope := []Coord{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	tail_positions := map[string]bool{rope[9].Hash(): true}

	for _, move := range moves {
		for ln := 0; ln < move.length; ln++ {
			rope[0].Move(move.dir)
			for i := 1; i < 10; i++ {
				rope[i].Adjust(rope[i-1])
			}
			tail_positions[rope[9].Hash()] = true
		}
	}

	return len(tail_positions)
}
