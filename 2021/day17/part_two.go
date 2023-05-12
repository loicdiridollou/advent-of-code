package main

func part2(input string) int {
	area := prepareData(input)
	count := 0

	for i := 1; i <= area.x_max; i++ {
		for j := area.y_min; j < 500; j++ {
			point := Point{0, 0}
			speed := Speed{i, j}
			max_y := area.y_min
			for area.canReach(point) {
				max_y = MaxInt(max_y, point.y)
				if area.contains(point) {
					count += 1
					break
				}
				point = point.update(speed)
				speed = speed.update()
			}
		}
	}

	return count
}
