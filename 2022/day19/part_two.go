package main

import (
	_ "embed"
	"os"
)

func part2() int {
	dat, _ := os.ReadFile("./day19-input")
  blueprints := parseInput(string(dat))
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
