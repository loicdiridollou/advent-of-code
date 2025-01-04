package main

import "strings"

// part2 function
func part2(input string, numr int, numc int) int {
	s := strings.Split(input, "\n")

	robots_vec := []Robot{}

	for _, line := range s {
		if line == "" {
			continue
		}
		robots_vec = append(robots_vec, parse_robot(line))
	}
	min_sf := numr * numc * numr * numc
	best_iter := 0

	for i := 1; i < numr*numc; i++ {
		quadrant_map := make(map[int]int)

		for j := 0; j < len(robots_vec); j++ {
			robots_vec[j] = robots_vec[j].move(1, numr, numc)

			quandrant := robots_vec[j].find_quadrant(numr, numc)
			if _, ok := quadrant_map[quandrant]; !ok {
				quadrant_map[quandrant] = 0
			}
			quadrant_map[quandrant]++
		}
		count := compute_score(quadrant_map)
		if count < min_sf {
			min_sf = count
			best_iter = i
		}
	}

	return best_iter
}
