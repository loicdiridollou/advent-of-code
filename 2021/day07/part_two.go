package main

func newDistance(v1, v2 int) int {
	var val int
	if v1-v2 < 0 {
		val = v2 - v1
	} else {
		val = v1 - v2
	}
	return val * (val + 1) / 2
}

func part2(input string) int {
	nums := prepareData(input)
	res := int(1e10)
	for i := 0; i < len(nums); i++ {
		if val := sum(Map(nums, newDistance, i)); val < res {
			res = val
		}
	}

	return res
}
