package main

func part2(input string) int {
	grid := prepareData(input)
	var num_flashes int

	for i := 1; i > 0; i++ {
		grid = addOne(grid)
		grid, num_flashes = flashOctopus(grid)

		if num_flashes == len(grid)*len(grid[0]) {
			return i
		}

	}

	return 0
}
