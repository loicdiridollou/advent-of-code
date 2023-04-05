package main

import (
	_ "embed"
)

func part2(input string) int {
	blueprints := parseInput(input)
	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	prod := 1
	for _, bp := range blueprints {
		st := newState(bp)
		geodesMade := st.calcMostGeodes(0, map[string]int{}, 32, 32)
		prod *= geodesMade
	}
	return prod
}
